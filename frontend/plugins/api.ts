export default defineNuxtPlugin((nuxtApp) => {
  const api = $fetch.create({
    baseURL: import.meta.server ? "http://caddy:8000/" : undefined,
    onRequest({ request, options, error }) {
      const acccessTokenCookie = useCookie("access_token");
      if (acccessTokenCookie.value) {
        options.headers.set(
          "Authorization",
          `Bearer ${acccessTokenCookie.value}`
        );
      }
    },
    async onResponseError({ response }) {
      if (response.status === 401) {
        await nuxtApp.runWithContext(() => navigateTo("/"));
      }
    },
  });

  return {
    provide: {
      api,
    },
  };
});
