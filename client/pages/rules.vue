<script lang="ts" setup>
const { data: page } = await useAsyncData("rules", () => {
  return queryCollection('rules').first()
})
</script>

<template>
    <div v-if="page" class="inhibit-defaults no-underline content-page p-8 m-8 text-left rounded-lg bg-primary/10 dark:bg-white/10 text-[var(--text-color)]">
        <ContentRenderer :value="page" class="rules" />
    </div>
</template>

<style scoped>

.rules a * {
    color: red;
    text-decoration: none;
}

</style>
