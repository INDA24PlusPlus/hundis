<template>
  <div>
    <div class="flex flex-col gap-4 mb-4">
      <UFormField label="Name">
        <UInput v-model="name" placeholder="My first contest!" />
      </UFormField>
      <UFormField label="Slug" description="Displayed in the URL">
        <UInput v-model="slug" placeholder="my-first-contest" />
      </UFormField>
      <UFormField label="Description">
        <UTextarea
          v-model="description"
          autoresize
          placeholder="The greatest contest ever!"
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
const contest = useContest();
const toast = useToast();

const props = defineProps({
  editId: {
    type: Number,
    default: null,
  },
});
const editMode = computed(() => !!props.editId);

const name = ref("");
const slug = ref("");
const description = ref("");
//const startTime = ref("");
//const endTime = ref("");

if (editMode.value) {
  const contestData = await contest.fetchContest(props.editId);
  if (!contestData) {
    navigateTo("/app/contests");
  } else {
    name.value = contestData.name;
    slug.value = contestData.slug;
    description.value = contestData.description;
    //startTime.value = contestData.startTime;
    //endTime.value = contestData.endTime;
  }
}

async function submitCreate() {
  if (editMode.value) return;

  try {
    const id = await contest.createContest({
      name: name.value,
      slug: slug.value,
      description: description.value,
      //startTime: startTime.value,
      //endTime: endTime.value,
    });

    toast.add({
      title: "Contest successfully created",
      color: "success",
    });

    navigateTo(`/app/contests/${id}`);
  } catch (error) {
    toast.add({
      title: "Could not create contest",
      description: error as string,
      color: "error",
    });
  }
}

async function submitEdit() {
  if (!props.editId) return;

  try {
    await contest.updateContest(props.editId, {
      name: name.value,
      slug: slug.value,
      description: description.value,
      //startTime: startTime.value,
      //endTime: endTime.value,
    });

    toast.add({
      title: "Contest successfully updated",
      color: "success",
    });
    navigateTo(`/app/contests/${props.editId}`);
  } catch (error) {
    toast.add({
      title: "Could not update contest",
      description: error as string,
      color: "error",
    });
  }
}

function cancelEdit() {
  navigateTo(`/app/contests/${props.editId}`);
}
</script>
