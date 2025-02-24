<script setup lang="ts">

</script>


<template>

<div class="nav-container">
    <nav>
        <div class="nav-item"><a href="#">Home</a></div>
        <div class="nav-item"><a href="#">About us</a></div>
        <div class="nav-item"><a href="#">Rules</a></div>
        <div class="nav-item"><a href="#">Explore</a></div>
    </nav>
</div>

</template>


<style scoped>

nav {
    display: flex;
    flex-direction: column;
    justify-content: center;
    gap: 5vh;
    align-items: center;
}

a {
    text-decoration: none;
    color: var(--text-color);
}

.nav-container {
    flex: 1;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: center;
    gap: 2vh;
}

</style>
