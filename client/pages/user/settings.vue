<script setup lang="ts">
const i18n = useI18n()
</script>

<template>
<div class="rules p-8 m-8 text-left rounded-lg bg-white/10 text-[var(--text-color)]">
    <div>
        <h1>{{ i18n.t('user_settings') }}</h1>
        <I18nSelect class="w-64" />
    </div>
</div>
</template>
