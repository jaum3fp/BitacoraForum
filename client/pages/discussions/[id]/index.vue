<script lang="ts" setup>
import { UAvatar, UTree } from '#components';
import { API } from '~/consts';
import { PostModel } from '~/models/post';
import type { PostModelType } from '~/models/post';


const route = useRoute()
const postId = parseInt(route.params.id as string)

const { data: post } = await useAsyncData<PostModelType>('getPost' + postId, async () => {
  const safeRoute = useRoute()
  const id = parseInt(safeRoute.params.id as string)
  await PostModel.incrementView(id)
  return await PostModel.getPostById(id)
})
const { data: comments, refresh: refreshComments } =
  await useAsyncData<Array<PostModelType>>('getComments' + postId, async () => await PostModel.getComments(postId))

const avatarImageUrl = computed(() =>
  post.value?.owner_avatar ? API.bitacoraForumAvatars + post.value.owner_avatar : 'https://avatars.githubusercontent.com/u/115469546?s=400&v=4')

</script>

<template>
<div v-if="post" class="block py-8 text-left">

    <div class="flex flex-wrap justify-between items-center">
        <header class="w-full">
            <h1 class="py-5 text-4xl">{{ post.title }}</h1>
        </header>
        <div class="flex items-center gap-5">
            <UAvatar size="3xl" :src="avatarImageUrl" />
            <p class="text-xl">By <ULink :to="'/user/' + post.owner_username">{{ post.owner_username }}</ULink></p>
        </div>
    </div>

    <div class="p-8 mt-4 text-left rounded-lg bg-white/10 text-[var(--text-color)] overflow-y-auto break-words my-5">
        {{ post.content }}
    </div>

    <div class="mb-10">
        <p>Created at: {{ post.created_at.toString().replace('T', ' ').replace('Z', '') }}</p>
        <p>Views: {{ post.views }}</p>
    </div>

    <h1 class="text-xl">Responses - {{ comments?.length }}</h1>

    <CommentsBox v-if="comments" :comments="comments" :parent="post" @send="refreshComments" />

</div>
</template>
