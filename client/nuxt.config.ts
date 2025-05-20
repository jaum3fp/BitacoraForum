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
    'pinia-plugin-persistedstate/nuxt'
  ],

  css: ['~/assets/css/global.css', '~/assets/css/tailwind.css'],

  svgo: {
    autoImportPath: '@/assets/svg',
    defaultImport: 'component',
    svgo: false
  },
})
