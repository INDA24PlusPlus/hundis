<template>
  <div class="overflow-x-auto">
    <table class="table-auto w-full border-collapse">
      <thead>
        <tr>
          <th v-for="column in columns" :key="column.key" class="px-4 py-2 cursor-pointer bg-sky-900 text-left"
            @click="column.sortable !== false && sortBy(column.key)">
            <div v-if="column.hidden !== true" class="flex items-center gap-2">

              {{ column.label }}
              <div class="flex flex-col text-xs">
                <span :class="[
                  'block',
                  sortKey === column.key && sortOrder === 'asc' ? 'text-white' : 'text-black/50'
                ]">
                  ▲
                </span>
                <span :class="[
                  'block',
                  sortKey === column.key && sortOrder === 'desc' ? 'text-white' : 'text-black/50'
                ]">
                  ▼
                </span>
              </div>
            </div>
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, index) in sortedItems" :key="item.id || item.slug" :class="[
          index % 2 === 0 ? 'bg-black/15' : 'bg-black/30',
          'hover:bg-black/45'
        ]">
          <td v-for="column in columns" :key="column.key" class="px-4 py-3">
            <template v-if="column.key === 'name' && item.href">
              <a :href="item.href" class="text-blue-500 hover:underline">
                {{ item[column.key] }}
              </a>
            </template>
            <template v-else>
              {{ item[column.key] }}
            </template>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from "vue";

// Define the structure of a column
interface Column {
  key: string;
  label: string;
  hidden?: boolean;
  sortable?: boolean;
}

// Define the structure of an item
interface Item {
  [key: string]: any;
}

// Props definition
const props = defineProps({
  items: {
    type: Array as () => Item[],
    required: true,
  },
  columns: {
    type: Array as () => Column[],
    required: true,
  },
});

// Reactive state for sorting
const sortKey = ref<string | null>(null);
const sortOrder = ref<"asc" | "desc">("asc");

// Method to handle sorting
function sortBy(key: string) {
  if (sortKey.value === key) {
    sortOrder.value = sortOrder.value === "asc" ? "desc" : "asc";
  } else {
    sortKey.value = key;
    sortOrder.value = "asc";
  }
}

// Computed property for sorted items
const sortedItems = computed<Item[]>(() => {
  if (!sortKey.value) return props.items;

  return [...props.items].sort((a, b) => {
    const aValue = a[sortKey.value as keyof Item];
    const bValue = b[sortKey.value as keyof Item];

    if (aValue < bValue) return sortOrder.value === "asc" ? -1 : 1;
    if (aValue > bValue) return sortOrder.value === "asc" ? 1 : -1;
    return 0;
  });
});
</script>