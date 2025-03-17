// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',

  devtools: { enabled: true },

  modules: ['nuxt-svgo', '@nuxt/ui'],

  css: ['~/assets/css/global.css', '~/assets/css/tailwind.css'],

  svgo: {
    autoImportPath: '@/assets/svg',
    global: false,
    defaultImport: 'component',
    svgo: false
  },
})