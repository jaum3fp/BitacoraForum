<script setup lang="ts"></script>

<template>

  <div class="flex w-full h-screen">
    <aside class="w-[200px] min-h-screen flex flex-col shadow-md">
      <Navbar />
    </aside>
    <main class="text-center px-12 flex-1 overflow-y-auto">
      <slot />
    </main>
  </div>

</template>
