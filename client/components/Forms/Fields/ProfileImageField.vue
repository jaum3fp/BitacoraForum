<script setup lang="ts">
import { API } from '~/consts'

const modelValue = defineModel<string | null | undefined>()

const fileInput = ref<HTMLInputElement | null>(null)

const triggerFileInput = () => fileInput.value?.click()

const onFileChange = (event: Event) => {
  const input = event.target as HTMLInputElement
  const file = input.files?.[0]
  if (file) {
    const reader = new FileReader()
    reader.onload = (e) => {
      modelValue.value = e.target?.result as string
    }
    reader.readAsDataURL(file)
  }
}

const avatarImageUrl = computed(() => {
  if (modelValue.value) {
    if (modelValue.value.includes(useUserStore().user?.username + ".")) {
      return API.bitacoraForumAvatars + modelValue.value
    } else {
      return modelValue.value
    }
  } else {
    return 'https://avatars.githubusercontent.com/u/115469546?s=400&v=4'
  }
})

</script>

<template>
  <div class="relative w-[120px] h-[120px] group cursor-pointer">

    <UAvatar :src="avatarImageUrl"  class="w-full h-full rounded-full">
      <template #fallback>
        <span class="w-full h-full flex items-center justify-center bg-gray-400 text-white text-sm font-medium rounded-full">
          JD
        </span>
      </template>
    </UAvatar>


    <div class="absolute inset-0 bg-black/50 rounded-full flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity z-10"
      @click="triggerFileInput"
    >
      <UIcon name="i-charm-camera" class="text-white text-2xl" />
    </div>

    <input ref="fileInput" type="file" accept="image/*" class="hidden" @change="onFileChange" />

  </div>
</template>
