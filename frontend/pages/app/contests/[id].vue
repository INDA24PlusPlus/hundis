<template>
  <div v-if="!isProblemRoute">
    <UCard>
      <div class="flex justify-between items-start mb-6">
        <h2 class="text-3xl">{{ contestData.name }}</h2>
        <UButton
          class="mt-1"
          @click="navigateTo(`/app/contests/${contestId}/edit`)"
        >
          Edit
        </UButton>
      </div>
      <ContestsNav />
      <NuxtPage />
    </UCard>
  </div>
  <div v-else>
    <NuxtPage />
  </div>
</template>

<script setup lang="ts">
const route = useRoute();
const contest = useContest();
const contestId = Number(route.params.id);

const contestData = ref({
  name: "",
  slug: "",
  description: "",
});

const fetchedContest = await contest.fetchContest(contestId);
if (!fetchedContest) {
  navigateTo("/app/contests");
} else {
  contestData.value = fetchedContest;
}

// Determine if the current route is a problem route
const isProblemRoute = computed(() =>
  route.path.startsWith(`/app/contests/${contestId}/problems/`)
  && route.path !== `/app/contests/${contestId}/problems/`
  && route.path !== `/app/contests/${contestId}/problems/add`
);
</script>