<script setup lang="ts">
import type { NavigationMenuItem } from '@nuxt/ui'


const items = ref<NavigationMenuItem[]>([
  {
    to: '/',
    icon: 'i-charm-home',
    label: 'Home',
  },
  {
    to: '/explore',
    icon: 'i-charm-map',
    label: 'Explore',
  },
  {
    to: '/rules',
    icon: 'i-charm-book',
    label: 'Rules',
  },
])

const active = ref()

defineShortcuts({
  1: () => active.value = '0',
  2: () => active.value = '1',
  3: () => active.value = '2',
})

</script>

<template>
  <div>
    <nav class="flex flex-col gap-[4vh] mt-4">
        <UNavigationMenu
            v-model="active"
            :items="items"
            :collapsed="false"
            orientation="vertical"
            class="data-[orientation=horizontal]:border-b border-default data-[orientation=horizontal]:w-full data-[orientation=vertical]:w-48"
        />
    </nav>
  </div>
</template>
