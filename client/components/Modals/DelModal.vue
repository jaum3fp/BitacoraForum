<script setup lang="ts">
const emit = defineEmits<{ close: [boolean] }>()
</script>

<template>
<UModal
    title="Eliminar description"
    description="¿Ha tomado la decisión de borrar el post? No lo podrás restaurar."
    :close="{ onClick: () => emit('close', false) }"
>
    <UButton icon="i-charm-bin" color="neutral" variant="outline"
    class="text-gray-600 hover:bg-red-600 transition-colors" />
    <template #footer>
        <div class="flex gap-2">
            <UButton color="error" label="Eliminar" @click="emit('close', true)" />
            <UButton color="neutral" label="Cancelar" variant="outline" @click="emit('close', false)" />
        </div>
    </template>
</UModal>
</template>
