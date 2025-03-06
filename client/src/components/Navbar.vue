<script setup lang="ts">

</script>


<template>

<div class="nav-container">
    <nav>
        <router-link to="/" class="nav-item">
            <i><svg  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round"  class="icon icon-tabler icons-tabler-outline icon-tabler-home"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M5 12l-2 0l9 -9l9 9l-2 0" /><path d="M5 12v7a2 2 0 0 0 2 2h10a2 2 0 0 0 2 -2v-7" /><path d="M9 21v-6a2 2 0 0 1 2 -2h2a2 2 0 0 1 2 2v6" /></svg></i>
            <p>Take a look</p>
        </router-link>
        <router-link to="/explore" class="nav-item">
            <i><svg  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round"  class="icon icon-tabler icons-tabler-outline icon-tabler-map"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M3 7l6 -3l6 3l6 -3v13l-6 3l-6 -3l-6 3v-13" /><path d="M9 4v13" /><path d="M15 7v13" /></svg></i>
            <p>Explore</p>
        </router-link>
        <router-link to="/rules" class="nav-item">
            <i><svg  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round"  class="icon icon-tabler icons-tabler-outline icon-tabler-book-2"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M19 4v16h-12a2 2 0 0 1 -2 -2v-12a2 2 0 0 1 2 -2h12z" /><path d="M19 16h-12a2 2 0 0 0 -2 2" /><path d="M9 8h6" /></svg></i>
            <p>Rules</p>
        </router-link>
    </nav>
</div>

</template>


<style scoped>

nav {
    display: flex;
    flex-direction: column;
    gap: 4vh;
    margin-top: 1em;
}

.nav-item {
    display: flex;
    gap: 1vh;
    align-items: center;
    padding: 0 1em;
}

.nav-item:hover {
    transition: background-color 0.6s ease;
    background-color: rgba(255, 255, 255, 0.1);
}

.nav-item:hover {
    transition: transform 0.3s ease;
    transform: translateX(10px);
}

.nav-item i {
    margin-right: 15px;
}

a {
    text-decoration: none;
    color: var(--text-color);
}

.nav-container {
    flex: 1;
    display: flex;
    flex-direction: column;
    gap: 2vh;
}

</style>
