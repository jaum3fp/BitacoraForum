<script setup lang="ts">
import { ref, onMounted } from 'vue';
import PostCard from './PostCard.vue';

const hello_world = ref("")

onMounted(async function() {
    const res = await fetch('http://127.21.0.12:8080')
    hello_world.value = await res.text()
})

</script>


<template>

<section class="posts-container">
    <PostCard title="Testing" :description="hello_world" author="root" image="none" />
</section>

</template>


<style scoped>

.posts-container {
    display: flex;
    flex-direction: column;
    gap: 3vh;
    margin: 6vh;
}

</style>
