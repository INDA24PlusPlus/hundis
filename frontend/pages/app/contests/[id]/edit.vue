<template>
  <div class="mt-8">
    <div class="px-4 py-6 bg-black/30">
      <h2 class="text-3xl mb-4">Edit contest</h2>
      <ContestsCreate :edit-id="Number($route.params.id)" />
    </div>
  </div>
</template>
<script lang="ts" setup></script>
