<template>
  <div class="flex flex-col">
    <div class="flex justify-between items-start">
      <h2 class="text-3xl mb-1">{{ problemData?.name }}</h2>
      <UButton
        class=""
        @click="navigateTo(`/app/problems/${problemSlug}/edit`)"
      >
        Edit
      </UButton>
    </div>
    <h4 class="text-slate-400 mb-4">{{ problemData?.slug }}</h4>
    <p>{{ problemData?.description }}</p>
  </div>
</template>

<script setup>
const route = useRoute();
const problem = useProblem();
const problemSlug = String(route.params.slug);

const problemData = ref({});
problem.fetchProblem(problemSlug).then(data => {
  problemData.value = data;
});
</script>
