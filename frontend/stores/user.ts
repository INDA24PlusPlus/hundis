import { defineStore, acceptHMRUpdate } from "pinia";

interface UserUpdateResponse {
  message: string;
}

export interface User {
  id: number | null;
  username: string;
  email: string;
  avatarUrl: string;
  permissions: string[];
}

export const useUser = defineStore("user", () => {
  const config = useRuntimeConfig();
  const accessTokenCookie = useCookie("access_token");

  // --- state ---
  const user = ref<User>({
    id: null,
    username: "",
    email: "",
    avatarUrl: "",
    permissions: [],
  });

  // --- getters ---
  const isLoggedIn = computed(() => {
    return user.value.id !== null;
  });

  // --- actions ---
  async function fetchUser() {
    if (accessTokenCookie.value && !isLoggedIn.value) {
      try {
        const data = await useAPI<User>("/api/me");
        user.value = data;
        return true;
      } catch {
        resetUser();
        return false;
      }
    }
  }

  async function login(next?: string) {
    if (!next) {
      const query = useRoute().query;
      next = query.next ? String(query.next) : "/app";
    }
    if (isLoggedIn.value) {
      return navigateTo(next);
    }
    return navigateTo(
      `https://github.com/login/oauth/authorize?client_id=${config.public.githubClientId}&redirect_uri=${window.origin}/api/auth/github&scope=read:user%20user:email&state=${next}`,
      { external: true }
    );
  }

  async function logout(next?: string) {
    if (!next) {
      const query = useRoute().query;
      next = query.next ? String(query.next) : "/";
    }
    accessTokenCookie.value = null;
    resetUser();
    return navigateTo(window.origin + next, { external: true });
  }

  async function updateAccount(username: string, email: string) {
    try {
      await useAPI<UserUpdateResponse>("/api/settings/account", {
        method: "PUT",
        body: { username, email },
      });
      user.value.username = username;
      user.value.email = email;
    } catch (error) {
      throw useAPIError(error);
    }
  }

  function resetUser() {
    user.value = {
      id: null,
      username: "",
      email: "",
      permissions: [],
      avatarUrl: "",
    };
  }

  return {
    user,
    isLoggedIn,
    login,
    logout,
    fetchUser,
    updateAccount,
  };
});

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useUser, import.meta.hot));
}
