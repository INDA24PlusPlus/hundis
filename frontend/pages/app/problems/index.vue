<template>
  <div>
    <UCard>
      <div class="mb-4 flex justify-between items-center">
        <h1 class="text-3xl">All Problems</h1>
        <div class="mb-4">
          <UButton @click="navigateTo('/app/problems/create')">Create Problem</UButton>
        </div>
      </div>
      <Table :items="problem.problems" :columns="columns" />
    </UCard>
  </div>
</template>


<script setup lang="ts">
import Table from "@/components/elements/table.vue";


const problem = useProblem();
await problem.fetchAllProblems();


problem.problems = problem.problems.map((problem) => {
  return {
    ...problem,
    name: problem.name,
    slug: problem.slug,
    description: problem.description,
    href: `/app/problems/${problem.slug}`,
  };
});


const columns = [
  { key: "name", label: "Name" },
  { key: "slug", label: "Slug" },
  { key: "description", label: "Description" },
];
</script>
