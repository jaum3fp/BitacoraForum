<script setup lang="ts">
import { ref, onMounted } from 'vue';
import PostCard from './PostCard.vue';

interface PostData {
    title: string
    content: string
    owner_username: string
}

let posts = ref<PostData[]>([])

onMounted(async function() {
    const res = await fetch('http://localhost:8080/api/v1/post/all')
    posts.value = await res.json()
})

</script>


<template>

<section class="d-flex flex-col gap-8 mx-[6vh]">
    <PostCard v-for="post in posts" 
        :title="post.title" 
        :description="post.content" 
        :author="post.owner_username" 
        image="none" />
</section>

</template>


<style scoped>
/*
.posts-container {
    display: flex;
    flex-direction: column;
    gap: 3vh;
    margin: 6vh;
}*/

</style>
