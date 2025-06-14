<script setup lang="ts">
import { UButton, UForm, UFormField, UTextarea } from '#components';
import type { FormSubmitEvent } from '@nuxt/ui';
import { PostModel, type NewPostModel, type PostModelType } from '~/models/post';
import { API } from '~/consts';


const props = defineProps<{
  comments: Array<PostModelType>,
  parent: PostModelType
}>()

const emit = defineEmits<{
  (e: 'send'): void
}>()

const userStore = useUserStore()

const state = {
  content: ''
} as const

type CommentState = typeof state

async function onSubmit(event: FormSubmitEvent<CommentState>) {
  if (userStore) {
    const comment: NewPostModel = {
      title: "Comment - " + props.parent.id,
      description: "Comment - " + props.parent.id,
      content: event.data.content,
      country_alpha: props.parent.country_alpha,
      owner_id: userStore.user?.id || 0,
      super_id: props.parent.id
    }
    const res = await PostModel.createPost(comment)
    if (res.success) {
      emit('send')
    }
  }
}

</script>

<template>

<div class="commentbox">
    <template v-for="comment in props.comments">
        <USeparator class="my-5" />
        <div class="commentcard">
            <div>
                <UAvatar size="md" :src="comment.owner_avatar ? API.bitacoraForumAvatars + comment.owner_avatar : 'https://avatars.githubusercontent.com/u/115469546?s=400&v=4'" />
                <ULink :to="'/user/' + comment.owner_username" class="ms-4">{{ comment.owner_username }}</ULink>
            </div>
            <div class="relative flex flex-col p-8 mt-4 text-left rounded-lg bg-white/10 text-[var(--text-color)] overflow-y-auto break-words">
                <div class="absolute top-2 right-2 text-sm text-gray-600 font-mono">
                    {{ comment.created_at.toString().replace('T', ' ').replace('Z', '') }}
                </div>
                {{ comment.content }}
            </div>
        </div>
    </template>

    <USeparator class="my-5" />

    <div v-if="userStore.user">
        <h1 class="text-xl">Add a comment:</h1>
        <UForm :state="state" class="my-5" @submit="onSubmit">
            <UFormField name="content" >
                <UTextarea v-model="state.content" class="w-full" />
            </UFormField>
            <UButton variant="solid" type="submit" class="mt-4">Send</UButton>
        </UForm>
    </div>
</div>

</template>
