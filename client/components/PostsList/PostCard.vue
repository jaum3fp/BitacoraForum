<script setup lang="ts">
import { UAvatar, UIcon } from '#components';
import { API } from '~/consts';


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
    author_avatar: string,
    views: number,
}>(), {
    modButtons: false
})

defineEmits<{
  (e: 'delete', id: number): void
}>()

const avatarUrl = props.author_avatar ? API.bitacoraForumAvatars + props.author_avatar : 'https://avatars.githubusercontent.com/u/115469546?v=4'

</script>

<template>

<div class="post-card flex bg-primary/4 via-transparent h-40 border-y-2 dark:border-2 border-y-secondary-300 dark:border-primary dark:rounded-md p-2 items-center gap-4" @click="navigateTo('/discussions/' + id)">
    <div class="post-card-header h-full flex items-center justify-center">
        <img class="h-full w-46 rounded" :src="props.flag" alt="Flag" />
    </div>
    <div class="post-card-content flex-1 flex flex-col justify-between h-full w-full overflow-y-hidden break-words">
        <div class="post-body relative">
            <div class="absolute top-0 left-0 text-dimmed flex items-center gap-2">
                <UIcon name="i-charm-eye" class="text-xl" />
                {{ props.views }}
            </div>
            <h2 class="text-lg font-semibold text-center w-full truncate mb-3">{{ props.title }}</h2>
            <p class="text-sm text-default/70 line-clamp-2 px-1 h-full">{{ props.description }}</p>
        </div>
        <div class="post-author font-bold text-xs text-left text-gray-500 px-1" @click.stop>
            <UAvatar :src="avatarUrl" size="sm" />
            <ULink :to="'/user/' + author" class="ms-2">{{ props.author }}</ULink>
        </div>
    </div>
    <div class="post-card-actions flex flex-col gap-2 items-start h-full px-2" @click.stop>
        <UChip position="bottom-right" size="3xl" :text="props.comments_total">
            <UButton icon="i-charm-messages" color="neutral" variant="outline" class="text-gray-600" @click="navigateTo('/discussions/' + id)" />
        </UChip>
        <UButton v-if="props.modButtons" icon="i-charm-bin" color="neutral" variant="outline"
            class="text-gray-600 hover:bg-red-600 transition-colors" @click="$emit('delete', props.id)" />
        <UButton v-if="props.modButtons" icon="i-charm-pencil" color="neutral" variant="outline"
            class="text-gray-600" @click="navigateTo('/discussions/' + id + '/edit')" />
    </div>
</div>

</template>
