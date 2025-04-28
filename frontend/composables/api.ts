import type { UseFetchOptions } from "nuxt/app";
import type { NitroFetchOptions, NitroFetchRequest } from "nitropack";

export function useAPIFetch<T>(
  url: string | (() => string),
  options?: UseFetchOptions<T>
) {
  return useFetch(url, {
    ...options,
    $fetch: useNuxtApp().$api as typeof $fetch,
  });
}

export function useAPI<T>(
  url: string,
  options?: NitroFetchOptions<NitroFetchRequest>
) {
  return $fetch<T>(url, {
    ...options,
  });
}

interface APIError {
  data: {
    error: string;
  };
}

export function useAPIError(error: unknown): string {
  return (error as APIError)?.data?.error ?? "An error occurred";
}
