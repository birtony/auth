/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

import { createRouter, createWebHistory } from 'vue-router';
import TheRoot from './TheRoot.vue';
import routes from './routes';
import supportedLocales from '@/config/supportedLocales';
import getStartingLocale from '@/mixins/i18n/getStartingLocale';
import { loadI18nMessages, updateI18nLocale } from '@/plugins/i18n';
import { i18n } from '../main';

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
  history: createWebHistory(getStartingLocale().base),
  routes: [
    // will match everything and put it under `$route.params.pathMatch`
    {
      path: '/:pathMatch(.*)*',
      name: 'NotFound',
      component: () => import(`../views/NotFound.vue`),
    },
    {
      path: `/:locale${getLocaleRegex()}?`,
      component: TheRoot,
      beforeEnter: (to, from) => {
        const locale = to.params.locale || getStartingLocale().base;
        const loadedLocales = i18n.global.availableLocales;
        console.log('loadedLocales', loadedLocales[0]);
        console.log('check', i18n.global.messages.value);
        // Load messages for starting locale
        if (
          loadedLocales.length === 1 &&
          !i18n.global.messages.value[loadedLocales[0]].length
        ) {
          loadI18nMessages(i18n.global, locale);
        }
        // If already loaded, check if a new locale was selected and load messages for it
        else {
          console.log('locale R', locale);
          updateI18nLocale(i18n.global, locale);
          // router.replace({
          //   name: to.params.name,
          //   params: {
          //     ...router.currentRoute._value.params,
          //     ...to.params,
          //     locale: locale,
          //   },
          //   query: to.query,
          // });
        }
        return;
      },
      children: routes,
    },
  ],
});

router.beforeEach((to, from, next) => {
  console.log('to', to);
  console.log('from', from);
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
