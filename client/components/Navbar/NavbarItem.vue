<script setup lang="ts">

type NavbarItemProps = {
  to: string;
  icon: string;
  text: string;
};

defineProps<NavbarItemProps>();
</script>

<template>
    <router-link
        :to="to"
        class="py-3 nav-item flex items-center gap-[1vh] px-4 transition-transform duration-300 hover:bg-white/10 hover:translate-x-2"
    >
        <UIcon :name="icon" class="size-6 mx-3" />
        <p>{{ text }}</p>
    </router-link>
</template>
