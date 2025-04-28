<template>
  <div class="mt-4 md:mt-8">
    <div class="mb-4 flex justify-between items-start">
      <h3 class="text-2xl">Problems</h3>
      <div>
        <UButton @click="navigateTo(`/app/contests/${contestId}/problems/add`)">Add a problem</UButton>
      </div>
    </div>
    <Table :items="problem.problems" :columns="columns" />
  </div>
</template>


<script setup lang="ts">
import Table from "@/components/elements/table.vue";

const problem = useProblem();
await problem.fetchAllProblems();

const route = useRoute();
const contestId = Number(route.params.id);

problem.problems = problem.problems.map((problem, i) => {
  return {
    ...problem,
    label: String.fromCharCode(65 + i),
    name: problem.name,
    slug: problem.slug,
    description: problem.description,
    href: `/app/contests/${contestId}/problems/${problem.slug}`,
  };
});

const columns = [
  { key: "label", label: "Label", hidden: true},
  { key: "name", label: "Name" },
  { key: "slug", label: "Slug" },
  { key: "description", label: "Description" },
];
</script>
