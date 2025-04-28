<template>
  <div>
    <nav class="bg-black/20 shadow-lg px-4 relative">
      <div class="flex items-center justify-between flex-wrap py-2">
        <div class="flex gap-x-8 items-center">
          <img
            src="~/assets/images/logo.png"
            alt=""
            width="48"
            class="rounded-lg"
          />
          <div class="font-bold text-xl">Hundis App</div>
        </div>
        <div>
          <span class="mr-4">{{ user.user.username }}</span>

          <UAvatar :src="user.user.avatarUrl" />
        </div>
      </div>
    </nav>
    <section class="flex">
      <div
        class="h-[calc(100vh-64px)] bg-black/35 px-2 pt-4 shadow-[4px_0_8px_rgba(0,0,0,0.2)] z-10"
      >
        <UNavigationMenu
          orientation="vertical"
          :items="items"
          class="data-[orientation=vertical]:w-52"
        />
      </div>
      <div class="grow p-4 background max-h-[calc(100vh-4rem)] overflow-y-auto">
        <slot />
      </div>
    </section>
  </div>
</template>
<script setup lang="ts">
import type { NavigationMenuItem } from "@nuxt/ui";

const user = useUser();

const items = ref<NavigationMenuItem[][]>([
  [
    {
      label: "Home",
      icon: "i-lucide-home",
      to: "/app",
    },
    {
      label: "Contests",
      icon: "i-lucide-circuit-board",
      defaultOpen: true,
      children: [
        {
          label: "Explore contests",
          icon: "i-lucide-telescope",
          to: "/app/contests",
        },
        {
          label: "Create contest",
          icon: "i-lucide-square-pen",
          to: "/app/contests/create",
        },
        {
          label: "View my contests",
          icon: "i-lucide-square-chart-gantt",
          disabled: true,
        },
      ],
    },
    {
      label: "Problems",
      icon: "i-lucide-layers",
      defaultOpen: true,
      children: [
        {
          label: "Browse problems",
          icon: "i-lucide-compass",
          to: "/app/problems",
          disabled: false,
        },
        {
          label: "Create problem",
          icon: "i-lucide-file-plus",
          to: "/app/problems/create",
          disabled: false,
        },
        {
          label: "View my problems",
          icon: "i-lucide-square-chart-gantt",
          disabled: true,
        },
      ],
    },
    {
      label: "Rankings",
      icon: "i-lucide-trophy",
      defaultOpen: true,
      children: [
        {
          label: "Global ranking",
          icon: "i-lucide-earth",
          disabled: true,
        },
        {
          label: "Country ranking",
          icon: "i-lucide-landmark",
          disabled: true,
        },
        {
          label: "University ranking",
          icon: "i-lucide-university",
          disabled: true,
        },
      ],
    },
    {
      label: "Settings",
      icon: "i-lucide-settings",
      defaultOpen: true,
      children: [
        {
          label: "Account",
          icon: "i-lucide-user",
          to: "/app/settings/account",
        },
      ],
    },
  ],
  [
    {
      label: "Logout",
      icon: "i-lucide-log-out",
      onSelect: () => {
        user.logout();
      },
    },
  ],
]);
</script>
<style scoped>
.background {
  background-image: linear-gradient(#141d31 0.6px, transparent 0.6px),
    linear-gradient(to right, #141d31 0.6px, #101727 0.6px);
  background-size: 16px 16px;
}
</style>
