// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },

  // ランタイム設定
  runtimeConfig: {
    public: {
      apiBase: 'http://localhost:8080'
    }
  },

  compatibilityDate: '2024-12-11'
})
