<template>
  <div>
    <div class="flex flex-col gap-4 mb-4">
      <UFormField label="Name">
        <UInput v-model="name" placeholder="My first problem!" />
      </UFormField>
      <UFormField label="Slug" description="Displayed in the URL">
        <UInput v-model="slug" placeholder="my-first-problem" />
      </UFormField>
      <UFormField label="Description">
        <UTextarea
          v-model="description"
          autoresize
          placeholder="The greatest problem ever!"
          type="textarea"
        />
      </UFormField>
      <!-- <UFormField label="Start time">
        <UInput v-model="startTime" type="datetime-local" />
      </UFormField>
      <UFormField label="End time">
        <UInput v-model="endTime" type="datetime-local" />
      </UFormField> -->
    </div>
    <div v-if="editMode" class="flex gap-4">
      <UButton @click="submitEdit">Edit</UButton>
      <UButton variant="outline" @click="cancelEdit">Cancel</UButton>
    </div>
    <UButton v-else @click="submitCreate">Create</UButton>
  </div>
</template>
<script lang="ts" setup>
const problem = useProblem();
const toast = useToast();

const props = defineProps({
  editSlug: {
    type: String,
    default: null,
  },
});
const editMode = computed(() => !!props.editSlug);

const name = ref("");
const slug = ref("");
const description = ref("");
//const startTime = ref("");
//const endTime = ref("");

if (editMode.value) {
  const problemData = await problem.fetchProblem(props.editSlug);
  if (!problemData) {
    navigateTo("/app/problems");
  } else {
    name.value = problemData.name;
    slug.value = problemData.slug;
    description.value = problemData.description;
    //startTime.value = problemData.startTime;
    //endTime.value = problemData.endTime;
  }
}

async function submitCreate() {
  if (editMode.value) return;

  try {
    const id = await problem.createProblem({
      name: name.value,
      slug: slug.value,
      description: description.value,
      author: "You",
      //startTime: startTime.value,
      //endTime: endTime.value,
    });

    toast.add({
      title: "Problem successfully created",
      color: "success",
    });

    navigateTo(`/app/problems/${id}`);
  } catch (error) {
    toast.add({
      title: "Could not create problem",
      description: error as string,
      color: "error",
    });
  }
}

async function submitEdit() {
  if (!props.editSlug) return;

  try {
    await problem.updateProblem(props.editSlug, {
      name: name.value,
      slug: slug.value,
      description: description.value,
      author: "You",
      //startTime: startTime.value,
      //endTime: endTime.value,
    });

    toast.add({
      title: "Problem successfully updated",
      color: "success",
    });
    navigateTo(`/app/problems/${props.editSlug}`);
  } catch (error) {
    toast.add({
      title: "Could not update problem",
      description: error as string,
      color: "error",
    });
  }
}

function cancelEdit() {
  navigateTo(`/app/problems/${props.editSlug}`);
}
</script>
