/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

import supportedLocales from '@/config/supportedLocales';
import getBrowserLocale from '@/mixins/i18n/getBrowserLocale';
import router from '@/router/index';

// Returns locale configuration. By default, try VUE_APP_I18N_LOCALE. As fallback, use 'en'.
function getMappedLocale(locale = import.meta.env.VUE_APP_I18N_LOCALE || 'en') {
  return supportedLocales.find((loc) => loc.id === locale);
}

export default function getStartingLocale() {
  // Get locale parameter form the URL
  const localeUrlParam = window.location.pathname
    .replaceAll(/^\//gi, '')
    .replace(/\/.*$/gi, '');
  console.log('localeUrlParam', localeUrlParam);
  // If locale parameter is set, check if it is amongst the supported locales and return it.
  if (
    localeUrlParam.length &&
    supportedLocales.find((loc) => loc.id === localeUrlParam)
  ) {
    console.log('found locale');
    return getMappedLocale(localeUrlParam);
  }
  // If no locale parameter is set in the URL, use the browser default.
  else {
    console.log('no locale');
    const browserLocale = getBrowserLocale({ countryCodeOnly: true });
    return getMappedLocale(browserLocale);
  }
}
