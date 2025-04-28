<template>
  <div>
    <UCard>


      <div class="mb-4 flex justify-between items-center">
        <h1 class="text-3xl">All Contests</h1>
        <div class="mb-4">
          <UButton class="" @click="navigateTo(`/app/contests/create`)">
            Create Contest
          </UButton>
        </div>
      </div>
      <Table :items="contest.contests" :columns="columns" />
    </UCard>
  </div>
</template>


<script setup lang="ts">
import Table from "@/components/elements/table.vue";


const contest = useContest();
await contest.fetchAllContests();


contest.contests = contest.contests.map((contest) => {
  return {
    ...contest,
    name: contest.name,
    slug: contest.slug,
    description: contest.description,
    href: `/app/contests/${contest.id}`,
  };
});


const columns = [
  { key: "name", label: "Name" },
  { key: "slug", label: "Slug" },
  { key: "description", label: "Description" },
];
</script>