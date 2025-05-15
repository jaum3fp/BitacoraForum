// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',

  devtools: { enabled: true },

  modules: ['nuxt-svgo', '@nuxt/ui', '@nuxt/content', '@pinia/nuxt'],

  css: ['~/assets/css/global.css', '~/assets/css/tailwind.css'],

  svgo: {
    autoImportPath: '@/assets/svg',
    defaultImport: 'component',
    svgo: false
  },
})
