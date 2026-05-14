export default defineNuxtConfig({
  modules: [
    '@nuxtjs/tailwindcss', 
    // '@nuxtjs/supabase'  <-- Comentalo por ahora
  ],
  
  nitro: {
    devProxy: {
      "/api": {
        target: "http://localhost:8080/api",
        changeOrigin: true,
        prependPath: true,
      },
    },
  },
})