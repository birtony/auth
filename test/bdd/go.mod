// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

module github.com/trustbloc/hub-auth/test/bdd

go 1.17

require (
	github.com/coreos/go-oidc/v3 v3.1.0
	github.com/cucumber/godog v0.9.0
	github.com/fsouza/go-dockerclient v1.6.5
	github.com/google/uuid v1.3.0
	github.com/ory/hydra-client-go v1.10.6
	github.com/pkg/errors v0.9.1
	github.com/tidwall/gjson v1.6.7
	github.com/trustbloc/edge-core v0.1.8-0.20220308160458-17fbc683162d
	github.com/trustbloc/hub-auth v0.0.0
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8
)

require (
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/Microsoft/go-winio v0.4.15-0.20200113171025-3fe6c5262873 // indirect
	github.com/Microsoft/hcsshim v0.8.9 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/asaskevich/govalidator v0.0.0-20200907205600-7a23bdc65eef // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/containerd/containerd v1.3.4 // indirect
	github.com/containerd/continuity v0.0.0-20200710164510-efbc4488d8fe // indirect
	github.com/cucumber/gherkin-go/v11 v11.0.0 // indirect
	github.com/cucumber/messages-go/v10 v10.0.3 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v1.4.2-0.20200319182547-c7ad2b866182 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/go-openapi/analysis v0.20.0 // indirect
	github.com/go-openapi/errors v0.20.1 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.5 // indirect
	github.com/go-openapi/loads v0.20.2 // indirect
	github.com/go-openapi/runtime v0.19.31 // indirect
	github.com/go-openapi/spec v0.20.3 // indirect
	github.com/go-openapi/strfmt v0.20.2 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-openapi/validate v0.20.2 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect
	github.com/hyperledger/aries-framework-go/spi v0.0.0-20220322085443-50e8f9bd208b // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.3 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/opencontainers/runc v0.1.1 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/sirupsen/logrus v1.6.0 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/tidwall/match v1.0.3 // indirect
	github.com/tidwall/pretty v1.0.2 // indirect
	go.mongodb.org/mongo-driver v1.5.1 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20211205182925-97ca703d548d // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa // indirect
	google.golang.org/grpc v1.42.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/square/go-jose.v2 v2.5.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	gotest.tools/v3 v3.0.3 // indirect
)

replace github.com/trustbloc/hub-auth => ../..
