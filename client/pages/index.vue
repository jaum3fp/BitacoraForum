<script setup lang="ts">
import SearchBar from '~/components/PostsList/SearchBar.vue';
import { UserModel } from '~/models/user';
</script>


<template>

    <header>
        <div class="py-10 text-center mx-[6vh]">
            <h1 class="text-4xl font-bold">Bitacora Forum</h1>
            <UButton @click="console.log(useUserStore().user)">Debug</UButton>
        </div>
    </header>
    <PostsList />

</template>


<style scoped>
</style>
