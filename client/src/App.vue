<script setup lang="ts">
import ProfileMenu from './components/ProfileMenu.vue';
import Navbar from './components/Navbar.vue';
</script>


<template>

  <aside>
    <ProfileMenu />
    <Navbar />
  </aside>
  <main>
    <router-view />
  </main>

</template>


<style scoped>

aside {
  width: 15vw;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

main {
  text-align: center;
  margin: 0 3em;
  width: 85%;
}

</style>
