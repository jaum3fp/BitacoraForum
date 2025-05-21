<script setup lang="ts">
import { UserModel } from '~/models/user';
import type { DropdownMenuItem } from '@nuxt/ui'
import UserProfileForm from './Forms/UserProfileForm.vue';


const userStore = useUserStore()

const showProfile = ref(false)


const onSelectLogout = async () => {
  const success = await UserModel.logout()
}

const items: DropdownMenuItem[][] = [
  [
    { label: 'Profile', icon: 'i-charm-id', onSelect() { showProfile.value = true } },
    { label: 'Settings', icon: 'i-charm-cog' },
  ],
  [{ label: 'Logout', color: 'error', icon: 'i-charm-sign-out', onSelect() { onSelectLogout() } } ]
]

</script>


<template>
    <div v-if="userStore.user" class="py-[6vh] flex justify-center">

        <UDropdownMenu
            arrow
            :items="items",
            :content="{ align: 'start', side: 'right', sideOffset: 30 }"
            :ui="{
                content: 'w-48',
                item: 'flex items-center gap-2',
                arrow: 'fill-current'
            }"
        >
            <UAvatar src="https://avatars.githubusercontent.com/u/115469546?v=4" size="3xl" />
        </UDropdownMenu>

    </div>
    <template v-else>
        <div class="flex flex-col gap-4 p-4">
            <UButton @click="navigateTo('auth/login')" class="h-12 rounded-md px-6" variant="outline" size="md">Log in</UButton>
            <UButton @click="navigateTo('auth/register')" class="h-12 rounded-md px-6" variant="outline" size="md">Regiser</UButton>
        </div>
    </template>

    <UserProfileForm v-model="showProfile" />

</template>
