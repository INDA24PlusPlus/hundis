<template>
  <div class="mt-8">
    <div class="px-4 py-6 bg-black/30">
      <h3 class="text-2xl">Welcome to {{ contestData.name }}!</h3>
      <hr class="mt-6 text-gray-600" />
      <p class="my-4">{{ contestData.description }}</p>
    </div>
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
</script>
