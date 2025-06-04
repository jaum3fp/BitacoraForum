<script setup lang="ts">
import { ref } from 'vue';
import PostCard from './PostCard.vue';
import SearchBar from './SearchBar.vue';
import { type PostsFilterType } from '~/models/post';
import PostsFilter from './PostsFilter.vue';
import { UButton } from '#components';
import DelModal from '@/components/Modals/DelModal.vue'
import { PostModel } from '~/models/post';

const props = withDefaults(defineProps<{
  filter?: PostsFilterType,
  mod?: boolean
}>(), {
  mod: false
})

defineEmits<{
  (e: 'add'): void
}>()

const countryStore = useCountriesStore()
const overlay = useOverlay()
const deleteModal = overlay.create(DelModal)

const searchFilter = ref(props.filter)

const { data: posts, refresh: refreshPosts } = useAsyncData(
  'posts' + props.filter,
  async () => await PostModel.getAllPosts(searchFilter.value|| {}),
  { server: false }
)

const onSearch = async (searchVal: string) => {
  const cc = countryStore.getAlphaByName(searchVal) || ''
  searchFilter.value = { ...props.filter, cc }
  refreshPosts()
}

const openDeleteModal = async (id: number) => {
  const deleteModalInstance = deleteModal.open()
  const toDelete = await deleteModalInstance.result
  if (toDelete) {
    const res = await PostModel.delete(id)
    if (res.success) {
      refreshPosts()
    }
  }
}

</script>

<template>

<div class="posts-listd-flex flex-col">
    <div class="posts-list-header mb-4 flex gap-4 w-full">
        <SearchBar @search="onSearch" :filter="{
          component: PostsFilter
        }" class="flex-1" />
        <UButton v-if="props.mod" icon="i-charm-plus" @click="$emit('add')">Add post</UButton>
    </div>
    <div class="posts-list-content">
        <PostCard v-if="posts" v-for="post in posts"
            :id="post.id"
            :title="post.title"
            :description="post.content"
            :country_alpha="post.country_alpha"
            :author="post.owner_username"
            :flag="post.flag"
            :mod-buttons="props.mod"
            @delete="openDeleteModal"
            image="none" />
        <USkeleton v-else />
    </div>
</div>

</template>
