/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

import { createRouter, createWebHistory } from 'vue-router';
import TheRoot from './TheRoot.vue';
import routes from './routes';
import supportedLocales from '@/config/supportedLocales';

// Creates regex (en|fr)
function getLocaleRegex() {
  let reg = '';
  supportedLocales.forEach((locale, index) => {
    reg = `${reg}${locale.id}${
      index !== supportedLocales.length - 1 ? '|' : ''
    }`;
  });
  return `(${reg})`;
}

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: `/:locale${getLocaleRegex()}?`,
      component: TheRoot,
      children: routes,
    },
  ],
});

router.beforeEach((to, from, next) => {
  // TODO: get locale dynamically
  const locale = 'en';
  if (to.params.locale && to.params.locale !== locale.id) {
    // router.replace({
    //   name: to.params.name,
    //   params: {
    //     ...router.currentRoute._value.params,
    //     ...to.params,
    //     locale: locale.base,
    //   },
    //   query: to.query,
    // });
    next();
    return;
  } else {
    next();
    return;
  }
});

export default router;
