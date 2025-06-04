<script lang="ts" setup>
import { PostModel } from '~/models/post';
import type { PostModelType } from '~/models/post';

const postId = parseInt(useRoute().params.id as string)
const { data: post } = await useAsyncData<PostModelType>('getPost' + postId, async () => {
  await PostModel.incrementView(postId.valueOf())
  return await PostModel.getPostById(postId)
})

</script>

<template>
<div v-if="post" class="block py-8 text-left">
    <header>
        <h1 class="py-5 text-4xl">{{ post.title }}</h1>
    </header>
    <p>Created at: {{ post.created_at.toString().replace('T', ' ').replace('Z', '') }}</p>
    <p>Views: {{ post.views }}</p>

    <div v-if="post" class="p-8 mt-4 text-left rounded-lg bg-white/10 text-[var(--text-color)]">

        {{ post.content }}

    </div>

</div>
</template>
