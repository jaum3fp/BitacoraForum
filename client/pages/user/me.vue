<script setup lang="ts">
import ProfileImageField from '~/components/Forms/Fields/ProfileImageField.vue';
import UserProfileForm from '~/components/Forms/UserProfileForm.vue';
import { API } from '~/consts';
import { UserModel, type UserModelType } from '~/models/user';


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

const onFormClose = () => {
  userRefresh()
}

watch(pimg, async () => await UserModel.updateUserData("profile_image", pimg.value))

const avatarImageUrl = computed(() => user.value?.profile_img ? API.bitacoraForumAvatars + user.value.profile_img : 'https://github.com/benjamincanac.png')

</script>

<template>

<div class="ownerProfile">
    <UserProfileLayout v-if="user">
        <template #avatar>
            <UAvatar :src="avatarImageUrl" class="w-[220px] h-[220px]" />
        </template>
        <template #user-data>
            <div class="w-full h-full flex items-center">
                <UButton icon="i-charm-pencil" variant="outline" class="absolute p-3 top-0 right-0" @click="showEditForm = true" />
                <h1 class="text-4xl font-bold">{{user.username}}'s page</h1>
                <h1 class="text-2xl">({{user.name}} {{user.surnames}})</h1>
            </div>
        </template>
        <template #body>
            <PostsList :filter="{ author: user.username }" />
        </template>
    </UserProfileLayout>
    <UserProfileForm v-model="showEditForm" @close="onFormClose" />
</div>

</template>
