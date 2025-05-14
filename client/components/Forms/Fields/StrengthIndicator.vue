<script setup lang="ts">

const modelValue = defineModel<string>()
const { color, score, text } = usePasswordStrength(modelValue)

</script>

<template>
    <div class="strengthindicator">
        <UProgress :model-value="score" :color="color" :indicator="text" :max="1" size="sm" />
        <p id="password-strength" class="text-sm font-medium">{{ text }}</p>
    </div>
</template>
