<script setup lang="ts">
import ProfileImageField from '~/components/Forms/Fields/ProfileImageField.vue';
import { UserModel, type UserModelType } from '~/models/user';


definePageMeta({
  validate: async (route) => {
    const user = await UserModel.getUserData(route.params.user as string)
    return user !== null
  }
})

const urlParam = useRoute().params.user as string

const { data: user } = useAsyncData('getUser', async () => await UserModel.getUserData(urlParam))

</script>

<template>

<div class="userpage">
    <UserProfileLayout v-if="user" :user="user">
        <template #avatar>
            <UAvatar src="https://github.com/benjamincanac.png" class="w-[220px] h-[220px]" />
        </template>
        <template #user-data>
            <h1 class="text-4xl font-bold">{{user.username}}</h1>
        </template>
        <template #body>
            <PostsList :filter="{ author: user.username }" />
        </template>
    </UserProfileLayout>
</div>

</template>
