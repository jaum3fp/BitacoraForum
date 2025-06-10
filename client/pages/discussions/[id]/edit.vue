<script lang="ts" setup>
import { PostModel } from '~/models/post';
import type { PostModelType } from '~/models/post';
import PostForm from '~/components/Forms/PostForm.vue';

const postId = parseInt(useRoute().params.id as string)
const { data: post } = await useAsyncData<PostModelType>('getPost' + postId, async () => await PostModel.getPostById(postId))

const showForm = ref(true)
watchEffect(() => showForm.value = true)

definePageMeta({
  middleware: ['validate-discussion-edit']
})

async function onSubmit(state: any) {
  const res = await PostModel.update(postId, {
    title: state.title,
    description: state.description,
    content: state.content,
    country_alpha: state.country_alpha,
  })
  if (res.success) {
    navigateTo('/user/me')
  }
}

</script>

<template>
<div v-if="post" class="block py-8 text-left">
    <PostForm v-model="showForm" :initial-data="post" title="Edit discussion" @close="navigateTo('/user/me')" @submit="onSubmit" />
</div>
</template>
