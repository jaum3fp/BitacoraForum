<script setup lang="ts">
import { ref, onMounted } from 'vue';
import PostCard from './PostCard.vue';
import { API } from '~/consts';
import { PostModel } from '~/models/post';
import { CountryModel } from '~/models/country';

const props = withDefaults(defineProps<{
  country?: string | null
}>(), {
  country: null
})

const { data: posts } = useAsyncData(
  'posts',
  async () => await PostModel.getAllPosts(props.country),
  { server: false }
)

</script>

<template>

<section class="d-flex flex-col gap-8 mx-[6vh]">
    <PostCard v-if="posts" v-for="post in posts"
        :title="post.title"
        :description="post.content"
        :country_alpha="post.country_alpha"
        :author="post.owner_username"
        :flag="post.flag"
        image="none" />
</section>

</template>
