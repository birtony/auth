/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

// Lazy load the component
function load(name) {
  return () => import(`../views/${name}.vue`);
}

export default [
  {
    path: '/sign-in',
    name: 'SignIn',
    component: load('SignIn'),
  },
  {
    path: '/sign-up',
    name: 'SignUp',
    component: load('SignUp'),
  },
];
