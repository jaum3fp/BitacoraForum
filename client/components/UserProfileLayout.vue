<template>
    <div class="grid grid-cols-8 grid-rows-3 gap-2 mx-[6vw] my-[4vh]">
      <div class="col-span-3 row-span-1 justify-center">
        <div class="col-start-2 flex justify-center items-center h-full">
            <slot name="avatar" />
        </div>
      </div>

      <div class="col-start-4 col-span-5 row-span-1">
          <div class="relative flex h-full w-full items-center">
              <slot name="user-data" />
          </div>
      </div>

      <div class="col-span-8 row-start-2 row-span-2 py-6">
          <slot name="body" />
      </div>

    </div>
</template>
