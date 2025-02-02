/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package authhandler

import (
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/trustbloc/auth/pkg/gnap/accesspolicy"
	"github.com/trustbloc/auth/pkg/gnap/api"
	"github.com/trustbloc/auth/pkg/gnap/session"
	"github.com/trustbloc/auth/spi/gnap"
)

/*
AuthHandler handles GNAP access requests and decides what access to grant.

TODO:
 - figure out how auth handler should work with login & consent handling

Input:
- The request handler passes the entire request object, since AuthHandler needs all parts of it:
  - The descriptors of requested access tokens and subject info are used to decide what access to give, and whether
    a login&consent interaction is necessary.
  - The client instance ID allows AuthHandler to check whether any requested
    tokens/data are already granted
  - The interact parameters allow the AuthHandler to decide which login&consent
    provider to use
- The request handler creates a request Verifier that will validate the client key bound to the request,
  and passes the Verifier into the AuthHandler.

TODO what AuthHandler needs to do:
 - Process request, use configured policy to:
   - deny forbidden requests
   - decide which requests can be granted based on access already saved in the session
   - decide which requests to collate into a login & consent interaction for user approval
 - If login & consent is necessary:
   - Decide which login & consent handler to use, invoke the handler to get the interact
     response, and respond to the client with the interact response
   - Handle the continue request, by fetching the login&consent result found under the given interact_ref,
     apply access policy to construct the tokens and subject data to return.
*/
type AuthHandler struct {
	continuePath string
	accessPolicy *accesspolicy.AccessPolicy
	sessionStore *session.Manager
	loginConsent api.InteractionHandler
}

// Config holds AuthHandler constructor configuration.
type Config struct {
	AccessPolicy       *accesspolicy.AccessPolicy
	ContinuePath       string
	InteractionHandler api.InteractionHandler
}

// New returns new AuthHandler.
func New(config *Config) *AuthHandler {
	return &AuthHandler{
		continuePath: config.ContinuePath,
		accessPolicy: config.AccessPolicy,
		sessionStore: session.New(),
		loginConsent: config.InteractionHandler,
	}
}

// HandleAccessRequest handles GNAP access requests.
func (h *AuthHandler) HandleAccessRequest( // nolint:funlen
	req *gnap.AuthRequest,
	reqVerifier api.Verifier,
) (*gnap.AuthResponse, error) {
	var (
		s   *session.Session
		err error
	)

	if req.Client == nil {
		// client can never be omitted entirely
		return nil, errors.New("missing client")
	}

	if req.Client.IsReference {
		s, err = h.sessionStore.GetByID(req.Client.Ref)
		if err != nil {
			return nil, fmt.Errorf("getting client session by client ID: %w", err)
		}
	} else {
		s, err = h.sessionStore.GetOrCreateByKey(req.Client.Key)
		if err != nil {
			return nil, fmt.Errorf("getting client session by key: %w", err)
		}
	}

	verifyErr := reqVerifier.Verify(s.ClientKey)
	if verifyErr != nil {
		return nil, fmt.Errorf("client request verification failure: %w", verifyErr)
	}

	permissions, err := h.accessPolicy.DeterminePermissions(req.AccessToken, s)
	if err != nil {
		return nil, fmt.Errorf("failed to determine permissions for access request: %w", err)
	}

	continueToken := gnap.AccessToken{
		Value: uuid.New().String(),
	}

	err = h.sessionStore.ContinueToken(&continueToken, s.ClientID)
	if err != nil {
		return nil, fmt.Errorf("saving continuation token to client session: %w", err)
	}

	err = h.sessionStore.SaveRequests(permissions.NeedsConsent, s.ClientID)
	if err != nil {
		return nil, fmt.Errorf("saving access requests to session: %w", err)
	}

	// TODO: figure out what parameters to pass into api.InteractionHandler.PrepareInteraction()
	// TODO: figure out where we save the client's finish redirect uri
	// TODO: support selecting one of multiple interaction handlers
	interact, err := h.loginConsent.PrepareInteraction(req.Interact)
	if err != nil {
		return nil, fmt.Errorf("creating response interaction parameters: %w", err)
	}

	resp := &gnap.AuthResponse{
		Continue: gnap.ResponseContinue{
			URI:         h.continuePath,
			AccessToken: continueToken,
		},
		Interact:   *interact,
		InstanceID: s.ClientID,
	}

	return resp, nil
}

// HandleContinueRequest handles GNAP continue requests.
func (h *AuthHandler) HandleContinueRequest(
	req *gnap.ContinueRequest,
	continueToken string,
	reqVerifier api.Verifier,
) (*gnap.AuthResponse, error) {
	s, err := h.sessionStore.GetByContinueToken(continueToken)
	if err != nil {
		return nil, fmt.Errorf("getting session for continue token: %w", err)
	}

	verifyErr := reqVerifier.Verify(s.ClientKey)
	if verifyErr != nil {
		return nil, fmt.Errorf("client request verification failure: %w", verifyErr)
	}

	consent, err := h.loginConsent.QueryInteraction(req.InteractRef)
	if err != nil {
		return nil, err
	}

	err = h.sessionStore.SaveSubjectData(consent.SubjectData, s.ClientID)
	if err != nil {
		return nil, err
	}

	newTokens := []gnap.AccessToken{}

	for _, tokenRequest := range consent.Tokens {
		tok := CreateToken(tokenRequest)

		newTokens = append(newTokens, *tok)

		err = h.sessionStore.AddToken(tok, s.ClientID)
		if err != nil {
			return nil, err
		}
	}

	resp := &gnap.AuthResponse{
		AccessToken: newTokens,
	}

	return resp, nil
}

// HandleIntrospection handles GNAP resource-server requests for access token introspection.
func (h *AuthHandler) HandleIntrospection( // nolint:gocyclo
	req *gnap.IntrospectRequest,
	reqVerifier api.Verifier,
) (*gnap.IntrospectResponse, error) {
	var (
		serverSession *session.Session
		clientSession *session.Session
		err           error
	)

	if req.ResourceServer == nil {
		return nil, errors.New("missing rs")
	}

	if req.ResourceServer.IsReference {
		serverSession, err = h.sessionStore.GetByID(req.ResourceServer.Ref)
		if err != nil {
			return nil, fmt.Errorf("getting rs session by rs ID: %w", err)
		}
	} else {
		// TODO: if we create a new session for an unfamiliar resource server, we're implicitly using a TOFU policy.
		serverSession, err = h.sessionStore.GetOrCreateByKey(req.ResourceServer.Key)
		if err != nil {
			return nil, fmt.Errorf("getting rs session by key: %w", err)
		}
	}

	verifyErr := reqVerifier.Verify(serverSession.ClientKey)
	if verifyErr != nil {
		return nil, fmt.Errorf("rs request verification failure: %w", verifyErr)
	}

	clientSession, clientToken, err := h.sessionStore.GetByAccessToken(req.AccessToken)
	if err != nil || clientToken == nil {
		return &gnap.IntrospectResponse{Active: false}, nil // nolint:nilerr
	}

	if req.Proof != "" && req.Proof != clientSession.ClientKey.Proof {
		return &gnap.IntrospectResponse{Active: false}, nil
	}

	subjectKeys := h.accessPolicy.AllowedSubjectKeys(clientToken.Access)

	resp := &gnap.IntrospectResponse{
		Active:      true,
		Access:      clientToken.Access,
		Key:         clientSession.ClientKey,
		Flags:       clientToken.Flags,
		SubjectData: map[string]string{},
	}

	for k := range subjectKeys {
		if v, ok := clientSession.SubjectData[k]; ok {
			resp.SubjectData[k] = v
		}
	}

	return resp, nil
}
