<script setup lang="ts">
import { ref } from 'vue';
import PostCard from './PostCard.vue';
import SearchBar from './SearchBar.vue';
import { type PostModelType, type PostsFilterType } from '~/models/post';
import PostsFilter from './PostsFilter.vue';
import { Paginator, UButton } from '#components';
import DelModal from '@/components/Modals/DelModal.vue'
import { PostModel } from '~/models/post';

interface PostList {
  filter?: PostsFilterType,
  mod?: boolean
}

const props = withDefaults(defineProps<PostList>(), {
  mod: false
})

defineEmits<{
  (e: 'add'): void
}>()

const i18n = useI18n()

const countryStore = useCountriesStore()
const overlay = useOverlay()
const deleteModal = overlay.create(DelModal)

const searchFilter = ref(props.filter)

const page = ref(1)
const itemsPerPage = ref(10)
const total = ref(0)

watch(() => page.value, () => refreshPosts())
watch(() => itemsPerPage.value, () => refreshPosts())

const { data: posts, refresh: refreshPosts } = useAsyncData<Array<PostModelType>>(
  'posts' + props.filter,
  async () => {
    if (!searchFilter.value) {
      searchFilter.value = {}
    }
    const start = itemsPerPage.value * (page.value - 1)
    const end = start + itemsPerPage.value
    searchFilter.value = { ...searchFilter.value, start, end }
    console.log(searchFilter.value)
    const res = await PostModel.getAllPosts(searchFilter.value)
    total.value = res.total
    return res.data
  },
  { server: false }
)

const onSearch = async (searchVal: string) => {
  const cc = countryStore.getAlphaByName(searchVal) || ''
  searchFilter.value = { ...props.filter, cc, title: searchVal }
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

<div v-if="posts" class="posts-listd-flex flex-col">
    <div class="posts-list-header mb-4 flex gap-4 w-full">
        <SearchBar @search="onSearch" :filter="{
          component: PostsFilter
        }" class="flex-1" />
        <UButton v-if="props.mod" icon="i-charm-plus" @click="$emit('add')">{{ i18n.t('user_me_add_post') }}</UButton>
    </div>
    <div class="posts-list-content">
        <template v-if="posts.length > 0">
            <Paginator v-model:page="page" v-model:items-per-page="itemsPerPage" :total="total" :sibling-count="3" />
            <PostCard v-for="post in posts"
                :id="post.id"
                :title="post.title"
                :description="post.description"
                :content="post.content"
                :country_alpha="post.country_alpha"
                :author="post.owner_username || ''"
                :flag="post.flag || ''"
                :mod-buttons="props.mod"
                :comments_total="post.comments_total"
                :author_avatar="post.owner_avatar || ''"
                :views="post.views"
                @delete="openDeleteModal"
                image="none" class="my-4" />
            <Paginator v-model:page="page" v-model:items-per-page="itemsPerPage" :total="total" :sibling-count="3" />
        </template>
        <USkeleton v-for="post in posts" v-else />
    </div>
</div>

</template>
