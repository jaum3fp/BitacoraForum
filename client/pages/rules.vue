<script lang="ts" setup>
const { data: page } = await useAsyncData("rules", () => {
  return queryCollection('rules').first()
})
</script>

<template>
    <div v-if="page" class="rules inhibit-defaults content-page p-8 m-8 text-left rounded-lg bg-white/10 text-[var(--text-color)]">
        <div>
            <ContentRenderer :value="page" />
        </div>
    </div>
</template>
