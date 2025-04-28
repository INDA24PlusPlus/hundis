<template>
  <div class="pt-8 px-4 md:px-6">
    <h2 class="text-3xl mb-4">Edit contest</h2>
    <ProblemsCreate :edit-slug="String($route.params.slug)" />
  </div>
</template>
<script lang="ts" setup></script>
