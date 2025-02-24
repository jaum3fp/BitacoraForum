<script setup lang="ts">
import ProfileMenu from './components/ProfileMenu.vue';
import Navbar from './components/Navbar.vue';
import PostsList from './components/PostsList.vue';
import SearchBar from './components/SearchBar.vue';
</script>


<template>

  <aside>
    <ProfileMenu />
    <Navbar />
  </aside>
  <main>
    <header>
      <div class="title-container">
        <h1>Bitacora Forum</h1>
      </div>
      <SearchBar />
    </header>
    <PostsList />
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

header {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 15vh;
  gap: 8vh;
}

main {
  text-align: center;
  margin: 0 3em;
  width: 85%;
}

</style>
