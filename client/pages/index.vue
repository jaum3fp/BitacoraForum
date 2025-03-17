<script setup lang="ts">
import SearchBar from '~/components/PostsList/SearchBar.vue';
</script>


<template>

    <header>
        <div class="py-10 text-center">
            <h1 class="text-4xl font-bold">Bitacora Forum</h1>
            <SearchBar />
        </div>
    </header>
    <PostsList />

</template>


<style scoped>
</style>
