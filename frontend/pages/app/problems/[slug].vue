<template>
  <div class="container p-0 my-0">
    <div class="flex p-0 my-0 bg-(--ui-bg) ring ring-(--ui-border) divide-y divide-(--ui-border) rounded-lg">      
      <div class="flex flex-col basis-1/2 border-r-1 border-slate-600">
        <ProblemsNav />
        <NuxtPage  />
      </div>
      <div class="flex flex-col basis-1/2">
        <ProblemsCode />
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
const route = useRoute();
const problem = useProblem();
const problemSlug = String(route.params.slug);

const problemData = ref<Problem>();

const fetchedProblem = await problem.fetchProblem(problemSlug);
if (!fetchedProblem) {
  navigateTo("/app/problems");
} else {
  problemData.value = fetchedProblem;
}
</script>
