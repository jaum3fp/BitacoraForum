// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',

  devtools: { enabled: true },

  devServer: {
    port: 3000,
    host: 'localhost',
    https: {
      key: '/certs/nuxt-app-key.pem',
      cert: '/certs/nuxt-app.pem',
    }
  },

  modules: [
    'nuxt-svgo',
    '@nuxt/ui',
    '@nuxt/content',
    '@pinia/nuxt',
    'pinia-plugin-persistedstate/nuxt',
    '@nuxtjs/i18n',
  ],

  css: ['~/assets/css/global.css', '~/assets/css/tailwind.css'],

  svgo: {
    autoImportPath: '@/assets/svg',
    defaultImport: 'component',
    svgo: false
  },

  i18n: {
    locales: [
      { code: 'es', language: 'es-ES', name: 'Español', file: 'es.json' },
      { code: 'en', language: 'en-US', name: 'English', file: 'en.json' },
    ],
    defaultLocale: 'es',
  }

})
