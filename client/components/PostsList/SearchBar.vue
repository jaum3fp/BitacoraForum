<script setup lang="ts">

import { ref, watchEffect } from 'vue';

const searchInput = ref("")

watchEffect(() => {
    console.log(searchInput.value)
})

</script>

<template>
    <div class="search-bar">
        <input v-model="searchInput" class="search-input" type="text" placeholder="Search" />
        <svg class="search-icon" xmlns="http://www.w3.org/2000/svg"  viewBox="0 0 24 24"><path stroke="none" d="M0 0h24v24H0z"/><path d="M10 10m-7 0a7 7 0 1 0 14 0a7 7 0 1 0 -14 0" /><path d="M21 21l-6 -6" /></svg>
    </div>
</template>

<style scoped>

.search-bar {
    display: flex;
    align-items: center;
    gap: 1em;

    padding: 1vh;
    padding-left: 1em;
    width: 30vw;
    border-radius: 2em;
    background-color: #f2f2f2;
}

.search-bar:hover {
    background-color: #e6e6e6;
}

.search-input {
    color: black;
    flex: 1;
    border: none;
    background-color: transparent;
    height: 100%;
}

.search-input:focus {
    outline: none;
}

.search-icon {
    color: white;
    border-radius: 100%;
    padding: 1vh;
    background-color: red;
    width: 24px;  
    height: 24px;
    fill: none;
    stroke: currentColor;
    stroke-width: 2px;
}

.search-icon:hover {
    cursor: pointer;
}

</style>