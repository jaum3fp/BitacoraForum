<script setup lang="ts">
import { CountryModel } from '~/models/country';
import { UserModel } from '~/models/user';


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
        <img class="h-full rounded" :src="flag" alt="Flag" />
    </div>
    <div class="post-card-content flex flex-col justify-between h-full w-full overflow-hidden">
        <div class="post-body">
            <h2 class="text-lg font-semibold text-center w-full truncate">{{ title }}</h2>
            <p class="text-sm text-gray-700 line-clamp-2 px-1">{{ description }}</p>
        </div>
        <div class="post-author text-xs text-gray-500 px-1">By {{ author }}</div>
    </div>
    <div class="post-card-actions flex items-start justify-center h-full">
        <UIcon name="i-charm-messages" class="size-6 text-gray-600" />
    </div>
</div>

</template>
