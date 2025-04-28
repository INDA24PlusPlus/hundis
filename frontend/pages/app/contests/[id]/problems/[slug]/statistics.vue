<template>
  <div class="pt-8 px-4 md:px-6">
    <h2 class="text-3xl mb-4">Statistics</h2>
    <div v-if="problemData?.author" class="mb-4">
      <h3 class="text-lg">Author:</h3>
      <div class="flex items-center mt-1">
        <p class="text-slate-300">{{ problemData?.author }}</p>
      </div>
    </div>
    <div v-if="problemData?.admin">
      <h3 class="text-lg">Admin</h3>
      <div class="flex items-center mt-1">
        <img
          :src="problemData?.admin?.avatarUrl"
          alt="Author Avatar"
          class="w-12 h-12 rounded-full mr-4"
        />
        <span class="text-slate-300">{{ problemData?.admin?.username }}</span>
      </div>
    </div>
    <div v-else>
      <p>Loading...</p>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useRoute } from 'vue-router';
const route = useRoute();
const problem = useProblem();
const problemSlug = String(route.params.slug);

const problemData = ref<Problem>();
problem.fetchProblem(problemSlug).then(data => {
  if(data){
    problemData.value = data;
  }
});
</script>