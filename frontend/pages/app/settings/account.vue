<template>
  <div>
    <UCard>
      <h2 class="text-3xl mb-4">Account settings</h2>
      <div class="flex flex-col gap-4 mb-4">
        <UFormField label="Username">
          <UInput v-model="username" placeholder="Enter your username" />
        </UFormField>
        <UFormField label="Email">
          <UInput v-model="email" placeholder="Enter your email" />
        </UFormField>
      </div>
      <UButton @click="save">Save</UButton>
    </UCard>
  </div>
</template>
<script lang="ts" setup>
const user = useUser();
const toast = useToast();

const username = ref(user.user.username);
const email = ref(user.user.email);

async function save() {
  try {
    await user.updateAccount(username.value, email.value);
    toast.add({
      title: "Account updated successfully",
      color: "success",
    });
  } catch (error) {
    toast.add({
      title: "Could not update account",
      description: error as string,
      color: "error",
    });
    return;
  }
}
</script>
