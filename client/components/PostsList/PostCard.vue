<script setup lang="ts">

const props = defineProps<{
    title: string,
    description: string,
    country_alpha: string,
    author: string,
    flag: string,
}>()
</script>


<template>

<div class="post-card flex h-40 border border-white p-2 items-center gap-4">
    <div class="post-card-header h-full flex items-center justify-center overflow-hidden">
        <img class="h-full rounded" :src="props.flag" alt="Flag" />
    </div>
    <div class="post-card-content flex flex-col justify-between h-full w-full overflow-hidden">
        <div class="post-body">
            <h2 class="text-lg font-semibold text-center w-full truncate">{{ props.title }}</h2>
            <p class="text-sm text-gray-700 line-clamp-2 px-1">{{ props.description }}</p>
        </div>
        <div class="post-author font-bold text-xs text-left text-gray-500 px-1">
            By <ULink :to="'/user/' + author">{{ props.author }}</ULink>
        </div>
    </div>
    <div class="post-card-actions flex items-start justify-center h-full px-2">
        <UChip position="bottom-right" size="3xl" :text="0">
            <UButton icon="i-charm-messages" color="neutral" variant="outline" class="text-gray-600" />
        </UChip>
    </div>
</div>

</template>
