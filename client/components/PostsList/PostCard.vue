<script setup lang="ts">

const props = withDefaults(defineProps<{
    id: number,
    title: string,
    description: string,
    content: string,
    country_alpha: string,
    author: string,
    flag: string,
    modButtons: boolean,
    comments_total: number,
}>(), {
    modButtons: false
})

defineEmits<{
  (e: 'delete', id: number): void
}>()

</script>

<template>

<div class="post-card flex h-40 border border-white p-2 items-center gap-4" @click="navigateTo('/discussions/' + id)">
    <div class="post-card-header h-full flex items-center justify-center overflow-hidden">
        <img class="h-full rounded" :src="props.flag" alt="Flag" />
    </div>
    <div class="post-card-content flex flex-col justify-between h-full w-full overflow-y-hidden break-words">
        <div class="post-body">
            <h2 class="text-lg font-semibold text-center w-full truncate">{{ props.title }}</h2>
            <p class="text-sm text-gray-700 line-clamp-2 px-1 h-full">{{ props.description }}</p>
        </div>
        <div class="post-author font-bold text-xs text-left text-gray-500 px-1" @click.stop>
            By <ULink :to="'/user/' + author">{{ props.author }}</ULink>
        </div>
    </div>
    <div class="post-card-actions flex flex-col gap-2 items-start h-full px-2" @click.stop>
        <UChip position="bottom-right" size="3xl" :text="props.comments_total">
            <UButton icon="i-charm-messages" color="neutral" variant="outline" class="text-gray-600" @click="navigateTo('/discussions/' + id)" />
        </UChip>
        <UButton v-if="props.modButtons" icon="i-charm-bin" color="neutral" variant="outline"
            class="text-gray-600 hover:bg-red-600 transition-colors" @click="$emit('delete', props.id)" @delete="console.log('del')" />
        <UButton v-if="props.modButtons" icon="i-charm-pencil" color="neutral" variant="outline"
            class="text-gray-600" @click="navigateTo('/discussions/' + id + '/edit')" />
    </div>
</div>

</template>
