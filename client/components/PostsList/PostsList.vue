<script setup lang="ts">
import { ref, onMounted } from 'vue';
import PostCard from './PostCard.vue';
import { API } from '~/consts';
import { PostModel } from '~/models/post';
import { CountryModel } from '~/models/country';
import SearchBar from './SearchBar.vue';
import { type PostsFilterType } from '~/models/post';
import PostsFilter from './PostsFilter.vue';


const props = defineProps<{
  filter?: PostsFilterType
}>()

const countryStore = useCountriesStore()

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

</script>

<template>

<div class="posts-listd-flex flex-col">
    <div class="posts-list-header mb-4">
        <SearchBar @search="onSearch" :filter="{
          component: PostsFilter
        }" />
    </div>
    <div class="posts-list-content">
        <PostCard v-if="posts" v-for="post in posts"
            :id="post.id"
            :title="post.title"
            :description="post.content"
            :country_alpha="post.country_alpha"
            :author="post.owner_username"
            :flag="post.flag"
            image="none" />
        <USkeleton v-else />
    </div>
</div>

</template>
