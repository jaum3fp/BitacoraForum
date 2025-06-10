<script lang="ts" setup>
const { locale } = useI18n()

const { data: pages } = await useAsyncData("rules" + locale.value, () => {
  return queryCollection('rules').all()
})

const page = computed(() => pages.value?.find(it => it.id.includes(`rules/index.${locale.value}.md`)))
</script>

<template>
    <div v-if="page" class="inhibit-defaults no-underline content-page p-8 m-8 text-left rounded-lg bg-primary/10 dark:bg-white/10 text-[var(--text-color)]">
        <ContentRenderer :value="page" class="rules" />
    </div>
</template>
