<script setup lang="ts">
import type { NavigationMenuItem } from '@nuxt/ui'
import { UserModel } from '~/models/user'
import { API } from '~/consts'


const userStore = useUserStore()

const onSelectLogout = async () => {
  const success = await UserModel.logout()
  if (success) navigateTo('/')
}

const buildItems = computed(() => {
  const items: NavigationMenuItem[][] = []
  const userItem: NavigationMenuItem = ({
    avatar: {
      src: userStore.user?.profile_img ? API.bitacoraForumAvatars + userStore.user.profile_img : 'https://avatars.githubusercontent.com/u/115469546?v=4'
    },
    label: userStore.user?.username,
    children: [
      {
        to: '/user/me',
        label: 'Profile',
        icon: 'i-charm-id'
      },
      {
        to: '/user/settings',
        label: 'Settings',
        icon: 'i-charm-cog'
      },
      {
        label: 'Logout',
        icon: 'i-charm-sign-out',
        class: 'text-red-800 fill-red',
        onSelect() { onSelectLogout() }
      },
    ],
  })
  const authItems: Array<NavigationMenuItem> = [
    {
      to: '/auth/login',
      icon: 'i-charm-key',
      label: 'Login',
    },
    {
      to: '/auth/register',
      icon: 'i-charm-notes',
      label: 'Register',
    },
  ]
  const staticItems: Array<NavigationMenuItem> = [
    {
      to: '/',
      icon: 'i-charm-home',
      label: 'Home',
    },
    {
      to: '/explore',
      icon: 'i-charm-map',
      label: 'Explore',
    },
    {
      to: '/rules',
      icon: 'i-charm-book',
      label: 'Rules',
    },
  ]
  items.push([], staticItems)
  items.push(userStore.user ? [userItem] : authItems)
  return items
})

const active = ref()
const collapse = ref(false)
watchEffect(() => console.log(active.value))
defineShortcuts({
  1: () => active.value = '0',
  2: () => active.value = '1',
  3: () => active.value = '2',
})

</script>

<template>
  <div>
    <nav class="flex flex-col gap-[4vh] mt-4">
        <UButton
            :icon="'i-charm-arrow-' + (collapse ? 'right' : 'left' )"
            class="ms-auto me-4"
            size="md"
            color="primary"
            variant="link"
            @click="collapse = !collapse"
        />
        <h1 class="text-xl text-center font-bold">Bitacora Forum</h1>
        <UNavigationMenu
            v-model="active"
            :items="buildItems"
            :collapsed="collapse"
            orientation="vertical"
            :ui="{
              link: 'gap-3 text-sm',
              item: '',
              list: 'space-y-6 my-4',
              root: 'w-full py-0'
            }"
        />
    </nav>
  </div>
</template>
