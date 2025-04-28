<template>
  <div>
    <nav class="bg-black/20 shadow-md px-3">
      <div
        class="container mx-auto flex items-center justify-between flex-wrap py-2"
      >
        <div class="flex gap-x-8 items-center">
          <div class="flex gap-x-6 items-center">
            <img
              src="~/assets/images/logo.png"
              alt=""
              width="48"
              class="rounded-lg"
            />
            <div class="font-bold text-xl">Hundis</div>
          </div>
          <NuxtLink
            to="/"
            class="hover:text-primary-400 transition-all duration-300 ease-in-out"
            >Home</NuxtLink
          >
          <NuxtLink
            class="hover:text-primary-400 transition-all duration-300 ease-in-out"
            to="/about"
            >About</NuxtLink
          >
          <NuxtLink
            class="hover:text-primary-400 transition-all duration-300 ease-in-out"
            to="/contact"
            >Contact</NuxtLink
          >
        </div>
        <div class="flex gap-x-2">
          <ClientOnly>
            <UButton
              v-if="user.isLoggedIn"
              variant="outline"
              @click="user.logout()"
              >Log out</UButton
            >
            <UButton @click="user.login()">{{
              user.isLoggedIn ? "Let's go" : "Log in"
            }}</UButton>

            <template #fallback>
              <UButton @click="user.login()">Log in</UButton>
            </template>
          </ClientOnly>
        </div>
      </div>
    </nav>
    <slot />
  </div>
</template>
<script setup lang="ts">
const user = useUser();
</script>
