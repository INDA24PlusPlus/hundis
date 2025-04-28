<template>
    <div class="flex h-18 border-b-1 border-slate-600 text-center items-center">
        <NuxtLink
        :to="`/app/contests/${contestId}/problems/${problemSlug}`"
        :class="{
            'basis-1/3 flex h-full pt-1 justify-center items-center px-4 border-b-2 border-blue-600 font-medium text-lg': isActive(`/app/contests/${contestId}/problems/${problemSlug}`),
            'basis-1/3 flex h-full pt-1 justify-center items-center px-4 border-b-2 border-transparent text-slate-400 font-medium text-lg': !isActive(`/app/contests/${contestId}/problems/${problemSlug}`)
        }"
        >
        <h3>Problem</h3>
        </NuxtLink>
        <NuxtLink
        :to="`/app/contests/${contestId}/problems/${problemSlug}/submissions`"
        :class="{
            'basis-1/3 flex h-full pt-1 justify-center items-center px-4 border-b-2 border-blue-600 font-medium text-lg': isActive(`/app/contests/${contestId}/problems/${problemSlug}/submissions`),
            'basis-1/3 flex h-full pt-1 justify-center items-center px-4 border-b-2 border-transparent text-slate-400 font-medium text-lg': !isActive(`/app/contests/${contestId}/problems/${problemSlug}/submissions`)
        }"
        >
        <h3>Submissions</h3>
        </NuxtLink>
        <NuxtLink
        :to="`/app/contests/${contestId}/problems/${problemSlug}/statistics`"
        :class="{
            'basis-1/3 flex h-full pt-1 justify-center items-center px-4 border-b-2 border-blue-600 font-medium text-lg': isActive(`/app/contests/${contestId}/problems/${problemSlug}/statistics`),
            'basis-1/3 flex h-full pt-1 justify-center items-center px-4 border-b-2 border-transparent text-slate-400 font-medium text-lg': !isActive(`/app/contests/${contestId}/problems/${problemSlug}/statistics`)
        }"
        >
        <h3>Statistics</h3>
        </NuxtLink>
    </div>
</template>

<script setup lang="ts">
const route = useRoute();
const problemSlug = String(route.params.slug);
const contestId = String(route.params.id);

const isActive = (path: string) => route.path == path;
</script>
