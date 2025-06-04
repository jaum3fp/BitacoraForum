<script lang="ts" setup>
import CreatePostForm from '~/components/Forms/CreatePostForm.vue';
import { PostModel } from '~/models/post';
import type { PostModelType } from '~/models/post';

const postId = parseInt(useRoute().params.id as string)
const { data: post } = await useAsyncData<PostModelType>('getPost' + postId, async () => {
  await PostModel.incrementView(postId.valueOf())
  return await PostModel.getPostById(postId)
})

definePageMeta({
  middleware: ['validate-discussion-edit']
})

</script>

<template>
<div v-if="post" class="block py-8 text-left">
    <CreatePostForm />
    {{post}}
</div>
</template>
