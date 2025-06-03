<template>
  <div class="flex flex-wrap gap-2 mx-[6vw] my-[4vh]">

    <div class="flex items-center w-full">

        <div class="w-3/8 flex justify-center items-center">
            <slot name="avatar" />
        </div>

        <div class="relative h-full flex-1 items-center">
            <slot name="user-data" />
        </div>

    </div>

    <div class="w-full py-6">
      <slot name="body" />
    </div>

  </div>
</template>
