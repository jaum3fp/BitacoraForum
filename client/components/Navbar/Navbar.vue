<script setup lang="ts">

import NavbarItem from './NavbarItem.vue';

const items = [
  { to: '/', icon: 'i-charm-home', text: 'Home' },
  { to: '/explore', icon: 'i-charm-map', text: 'Explore' },
  { to: '/rules', icon: 'i-charm-book', text: 'Rules' },
];
</script>

<template>
  <div>
    <nav class="flex flex-col gap-[4vh] mt-4">
      <template v-for="item in items">
        <NavbarItem :to="item.to" :icon="item.icon" :text="item.text" />
      </template>
    </nav>
  </div>
</template>
