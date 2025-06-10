<script setup lang="ts"></script>

<template>

  <div class="flex w-full h-screen">
    <aside class="min-w-[12%] min-h-screen flex flex-col shadow-md">
      <Navbar />
    </aside>
    <main class="text-center px-12 flex-1 overflow-y-auto">
      <div class="my-8">
        <NavBreadcrumb />
      </div>
      <slot />
    </main>
  </div>

</template>
