<script setup lang="ts">
import { PostsList } from '#components';
import PostForm from '~/components/Forms/PostForm.vue';
import ProfileImageField from '~/components/Forms/Fields/ProfileImageField.vue';
import UserProfileForm from '~/components/Forms/UserProfileForm.vue';
import { API } from '~/consts';
import { UserModel, type UserModelType } from '~/models/user';
import { PostModel } from '~/models/post';


definePageMeta({
  validate: async (route) => {
    const user = await UserModel.getOwnData()
    return user !== null
  }
})

const userStore = useUserStore()

const { data: user, refresh: userRefresh } = useAsyncData('getUser', async () => await userStore.loadAndGet())

const pimg = ref("")
const showEditForm = ref(false)
const showCreateForm = ref(false)

watch(pimg, async () => await UserModel.updateUserData("profile_image", pimg.value))

const avatarImageUrl = computed(() => user.value?.profile_img ? API.bitacoraForumAvatars + user.value.profile_img : 'https://avatars.githubusercontent.com/u/115469546?s=400&v=4')

async function onSubmit(state: any) {
  if (userStore.user) {
    const res = await PostModel.createPost({
      title: state.title,
      description: state.description,
      content: state.content,
      country_alpha: state.country_alpha,
      owner_id: userStore.user.id,
    })
    if (res.success) {
      refreshNuxtData('posts' + { author: user.value?.username })
      showCreateForm.value = false
    }
  }
}

</script>

<template>

<div class="ownerProfile">
    <UserProfileLayout v-if="user">
        <template #avatar>
            <UAvatar :src="avatarImageUrl" class="w-[220px] h-[220px]" />
        </template>
        <template #user-data>
            <div class="w-full h-full flex gap-4 items-center">
                <UButton icon="i-charm-pencil" variant="outline" class="absolute p-3 top-0 right-0" @click="showEditForm = true" />
                <h1 class="text-4xl font-bold">{{user.username}}'s page</h1>
                <h1 class="text-2xl">({{user.name}} {{user.surnames}})</h1>
            </div>
        </template>
        <template #body>
            <PostsList :filter="{ author: user.username }" :mod="true" @add="showCreateForm = true" />
        </template>
    </UserProfileLayout>
    <UserProfileForm v-model="showEditForm" @close="userRefresh()" />
    <PostForm v-model="showCreateForm" title="Create new discussion" @submit="onSubmit" />
</div>

</template>
