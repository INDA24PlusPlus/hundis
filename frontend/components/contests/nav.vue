<template>
    <div class="flex h-12 border-b-1 border-slate-600 text-center items-center">
        <NuxtLink
        :to="`/app/contests/${contestId}`"
        :class="{
            'flex h-full pt-1 justify-center items-center px-8 border-b-2 border-blue-600 font-medium text-lg': isActive('/app/contests/' + contestId),
            'flex h-full pt-1 justify-center items-center px-8 border-b-2 border-transparent text-slate-400 font-medium text-lg': !isActive('/app/contests/' + contestId)
        }"
        >
        <h3>Contest</h3>
        </NuxtLink>
        <NuxtLink
        :to="`/app/contests/${contestId}/problems`"
        :class="{
            'flex h-full pt-1 justify-center items-center px-8 border-b-2 border-blue-600 font-medium text-lg': isActive('/app/contests/' + contestId + '/problems'),
            'flex h-full pt-1 justify-center items-center px-8 border-b-2 border-transparent text-slate-400 font-medium text-lg': !isActive('/app/contests/' + contestId + '/problems')
        }"
        >
        <h3>Problems</h3>
        </NuxtLink>
        <NuxtLink
        :to="`/app/contests/${contestId}/standings`"
        :class="{
            'flex h-full pt-1 justify-center items-center px-8 border-b-2 border-blue-600 font-medium text-lg': isActive('/app/contests/' + contestId + '/standings'),
            'flex h-full pt-1 justify-center items-center px-8 border-b-2 border-transparent text-slate-400 font-medium text-lg': !isActive('/app/contests/' + contestId + '/standings')
        }"
        >
        <h3>Standings</h3>
        </NuxtLink>
    </div>
</template>

<script setup lang="ts">
const route = useRoute();
const contestId = String(route.params.id);

const isActive = (path: string) => route.path === path;
</script>
