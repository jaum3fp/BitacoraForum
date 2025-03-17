<script setup lang="ts">
const items = ref([
    { label: 'Profile', icon: 'i-chamber-user' },
    { label: 'Settings', icon: 'i-chamber-cog' },
    { label: 'Theme', icon: 'i-chamber-cog', children: [
      { label: 'Light', icon: 'i-chamber-sun' },
      { label: 'Dark', icon: 'i-chamber-moon' },
      { label: 'Auto', icon: 'i-chamber-auto' }
    ] },
    { label: 'Logout', icon: 'i-chamber-logout' }
])
</script>


<template>
  <div class="py-[6vh] flex justify-center">
    <UDropdownMenu
      arrow
      :items="items",
      :content="{
        align: 'start',
        side: 'right',
        sideOffset: 30
      }"
      :ui="{
        content: 'w-48 bg-red-500',
        item: 'flex items-center gap-2',
        arrow: 'fill-current'
      }"  
    >
      <UAvatar src="https://avatars.githubusercontent.com/u/115469546?v=4" size="3xl" />
    </UDropdownMenu>
  </div>
</template>
