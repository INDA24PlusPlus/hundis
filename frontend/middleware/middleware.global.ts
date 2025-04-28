export default defineNuxtRouteMiddleware(async (to, from) => {
  const user = useUser();
  await user.fetchUser();
});
