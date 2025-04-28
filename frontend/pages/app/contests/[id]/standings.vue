<template>
    <div class="mt-4 md:mt-8">
        <div class="mb-4 flex justify-between items-center">
            <h3 class="text-2xl">Standings</h3>
        </div>
        <Table :items="standings" :columns="columns" />
    </div>
</template>

<script setup lang="ts">
import Table from "@/components/elements/table.vue";

const standings = [
    {
        name: "Alice",
        solved: 3,
        time: "1:30",
        problems: { A: "Pass", B: "Pass", C: "Pass", D: "Fail" },
    },
    {
        name: "Bob",
        solved: 2,
        time: "2:15",
        problems: { A: "Pass", B: "Fail", C: "Pass", D: "Fail" },
    },
    {
        name: "Charlie",
        solved: 4,
        time: "1:10",
        problems: { A: "Pass", B: "Pass", C: "Pass", D: "Pass" },
    },
    {
        name: "David",
        solved: 1,
        time: "3:00",
        problems: { A: "Fail", B: "Fail", C: "Pass", D: "Fail" },
    },
];

const columns = [
    { key: "name", label: "Name" },
    { key: "solved", label: "Solved" },
    { key: "time", label: "Time" },
    { key: "problems.A", label: "A" },
    { key: "problems.B", label: "B" },
    { key: "problems.C", label: "C" },
    { key: "problems.D", label: "D" },
];
</script>