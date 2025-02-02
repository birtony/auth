<!--
 * Copyright SecureKey Technologies Inc. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
-->

<script setup>
import TheToastNotification from '@/components/TheToastNotification.vue';
import IconLogo from '@/components/icons/IconLogo.vue';
import IconSpinner from '@/components/icons/IconSpinner.vue';
import useBreakpoints from '@/plugins/breakpoints.js';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
</script>

<template>
  <the-toast-notification
    v-if="systemError"
    :title="t('SignUp.errorToast.title')"
    :description="t('SignUp.errorToast.description')"
    type="error"
  />
  <div
    class="overflow-hidden h-auto text-xl bg-gradient-dark rounded-xl md:max-w-4xl md:text-3xl"
  >
    <div
      class="grid grid-cols-1 w-full h-full bg-no-repeat divide-x divide-neutrals-medium md:grid-cols-2 md:px-20 bg-onboarding-flare-lg divide-opacity-25"
    >
      <div class="hidden col-span-1 py-24 pr-16 md:block">
        <IconLogo class="mb-12" />

        <div class="flex overflow-y-auto flex-1 items-center mb-8 max-w-full">
          <img
            class="flex w-10 h-10"
            src="@/assets/signup/onboarding-icon-1.svg"
          />
          <span class="pl-5 text-base text-neutrals-white align-middle">
            {{ t('SignUp.leftContainer.span1') }}
          </span>
        </div>

        <div class="flex overflow-y-auto flex-1 items-center mb-8 max-w-full">
          <img
            class="flex w-10 h-10"
            src="@/assets/signup/onboarding-icon-2.svg"
          />
          <span class="pl-5 text-base text-neutrals-white align-middle">
            {{ t('SignUp.leftContainer.span2') }}
          </span>
        </div>

        <div class="flex overflow-y-auto flex-1 items-center max-w-full">
          <img
            class="flex w-10 h-10"
            src="@/assets/signup/onboarding-icon-3.svg"
          />
          <span class="pl-5 text-base text-neutrals-white align-middle">
            {{ t('SignUp.leftContainer.span3') }}
          </span>
        </div>
      </div>
      <div class="object-none object-center col-span-1 md:block">
        <div class="px-6 md:pt-16 md:pr-0 md:pb-12 md:pl-16">
          <IconLogo class="justify-center my-2 mt-12 md:hidden" />
          <div class="items-center pb-6 text-center">
            <h1 class="text-2xl font-bold text-neutrals-white md:text-4xl">
              {{ t('SignUp.heading') }}
            </h1>
          </div>
          <div
            class="grid grid-cols-1 gap-5 justify-items-center content-center mb-8 w-full h-64"
          >
            <IconSpinner v-if="loading" />
            <button
              v-for="(provider, index) in providers"
              v-else
              :id="provider.id"
              :key="index"
              class="flex flex-wrap items-center w-full h-11 text-sm font-bold text-neutrals-dark bg-neutrals-softWhite rounded-md"
              @click="beginOIDCLogin(provider.id)"
              @keyup.enter="beginOIDCLogin(provider.id)"
            >
              <img :src="provider.SignUpLogoUrl" />
            </button>
          </div>
          <div class="mb-8 text-center">
            <p class="text-base font-normal text-neutrals-white">
              {{ t('SignUp.redirect') }}
              <router-link
                class="text-primary-blue whitespace-nowrap underline-blue"
                :to="{ name: 'SignIn' }"
                >{{ t('SignUp.signin') }}</router-link
              >
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      providers: [],
      statusMsg: '',
      loading: true,
      systemError: false,
      breakpoints: useBreakpoints(),
    };
  },
};
</script>
