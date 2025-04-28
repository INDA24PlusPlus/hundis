// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: "2024-11-01",
  devtools: { enabled: true },
  modules: [
    "@nuxt/ui",
    "@nuxt/eslint",
    "@pinia/nuxt",
    "@nuxt/icon",
    (_options, nuxt) => {
      nuxt.hooks.hook("vite:extendConfig", (config) => {
        config.server = {
          ...config.server,
          hmr: {
            port: 3333,
            clientPort: 3333,
          },
        };
      });
    },
  ],
  css: ["~/assets/css/main.css"],
  ui: {
    theme: {
      colors: ["primary", "secondary", "success", "info", "warning", "error"],
    },
  },
  watchers: {
    chokidar: {
      usePolling: process.env.DOCKER_DEV === "true",
    },
  },
  vite: {
    server: {
      hmr: {
        port: 3333,
        clientPort: 3333,
      },
    },
  },
  runtimeConfig: {
    public: {
      githubClientId:
        process.env.NODE_ENV == "development"
          ? "Ov23liLUIbrO4iuhaVwF"
          : "Ov23li1H2T5Ig5IKOybl",
    },
  },
  colorMode: {
    preference: "dark",
  },
  routeRules: {
    "/app/**": {
      ssr: false,
      appMiddleware: ["auth"],
    },
    "/api/**": {
      proxy: "http://caddy:8000/api",
    },
  },
  icon: {
    localApiEndpoint: "/nuxt_api/_nuxt_icon",
    provider: "server",
    collections: ["lucide", "simple-icons"],
  },
});
