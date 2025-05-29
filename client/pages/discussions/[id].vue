<script lang="ts" setup>
import { PostModel } from '~/models/post';
import type { PostModelType } from '~/models/post';

const postId = parseInt(useRoute().params.id as string)
const { data: post } = await useAsyncData<PostModelType>('getPost' + postId, async () => await PostModel.getPostById(postId))



</script>

<template>
<div v-if="post">
    <header>
        <h1>{{ post.title }}</h1>
    </header>
    <p>Created at: {{ post.created_at }}</p>
    <p>Modified at: {{ post.updated_at }}</p>

    <div v-if="post" class="p-8 m-8 text-left rounded-lg bg-white/10 text-[var(--text-color)]">

        {{ post.content }}

    </div>

</div>
</template>
