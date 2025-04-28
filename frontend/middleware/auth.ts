// Will only be called for /app routes
export default defineNuxtRouteMiddleware(async (to, from) => {
  const user = useUser();

  if (!user.isLoggedIn) {
    return navigateTo("/");
  }

  setPageLayout("app");
});
