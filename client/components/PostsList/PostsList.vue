<script setup lang="ts">
import { ref, onMounted } from 'vue';
import PostCard from './PostCard.vue';
import { API } from '~/consts';
import { PostModel } from '~/models/post';

const props = defineProps<{
  country?: string
}>()

interface PostData {
    title: string
    content: string
    owner_username: string
}

const { data: posts } = useAsyncData('getAllPosts', async () => await PostModel.getAllPosts(props.country && props.country))

</script>


<template>

<section class="d-flex flex-col gap-8 mx-[6vh]">
    <PostCard v-for="post in posts"
        :title="post.title"
        :description="post.content"
        :country_alpha="post.country_alpha"
        :author="post.owner_username"
        image="none" />
</section>

</template>
