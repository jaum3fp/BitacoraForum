<script setup lang="ts">

const modelValue = defineModel<string>()
const show = ref(false)

const { color, score, text } = usePasswordStrength(modelValue)


</script>

<template>
  <div class="space-y-2">
    <UInput
        v-model="modelValue"
        :placeholder="$attrs.placeholder as string || 'Password'"
        :color="color || 'neutral'"
        :type="show ? 'text' : 'password'"
        :ui="($attrs.ui as object) || { trailing: 'pe-1' }"
        aria-describedby="password-strength"
        class="w-full"
    >
        <template #trailing>
            <UButton
                color="neutral"
                variant="link"
                size="sm"
                :icon="show ? 'i-lucide-eye-off' : 'i-lucide-eye'"
                :aria-label="show ? 'Hide password' : 'Show password'"
                :aria-pressed="show"
                aria-controls="password"
                @click="show = !show"
            />
        </template>
    </UInput>

  </div>
</template>
