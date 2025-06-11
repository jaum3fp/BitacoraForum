<script setup lang="ts">
import type { NavigationMenuItem } from '@nuxt/ui'
import { UserModel } from '~/models/user'
import { API } from '~/consts'


const userStore = useUserStore()
const countriesStore = useCountriesStore()
const i18n = useI18n()
const localePath = useLocalePath()
const switchLocalePath = useSwitchLocalePath()

const onSelectLogout = async () => {
  const success = await UserModel.logout()
  if (success) navigateTo('/')
}

const onSelectLang = (cc: any) => {
  navigateTo(switchLocalePath(cc))
}

const buildItems = computed(() => {
  const items: NavigationMenuItem[][] = []
  const i18nItemChildren = i18n.locales.value.map(lang => {
    const cc = lang.language?.split('-')[1]
    return {
      label: lang.name,
      value: lang.code,
      avatar: { src: cc ? countriesStore.getCountryByAlpha(cc.toUpperCase()).flagSvg : '' },
      onSelect() { onSelectLang(lang.code) }
    }
  })
  const i18nItem: NavigationMenuItem = ({
    avatar: {
      src: i18nItemChildren.find(item => item.value === i18n.locale.value)?.avatar.src
    },
    label: i18n.locales.value.find(it => it.code === i18n.locale.value)?.name,
    children: i18nItemChildren
  })
  const userItem: NavigationMenuItem = ({
    avatar: {
      src: userStore.user?.profile_img ?
      API.bitacoraForumAvatars + userStore.user.profile_img :
      'https://avatars.githubusercontent.com/u/115469546?v=4'
    },
    label: userStore.user?.username,
    children: [
      {
        to: localePath('/user/me'),
        label: i18n.t("nav_user_profile"),
        icon: 'i-charm-id'
      },
      {
        to: localePath('/user/settings'),
        label: i18n.t('nav_user_settings'),
        icon: 'i-charm-cog'
      },
      {
        label: i18n.t('nav_user_logout'),
        icon: 'i-charm-sign-out',
        class: 'text-red-800 fill-red',
        onSelect() { onSelectLogout() }
      },
    ],
  })
  const authItems: Array<NavigationMenuItem> = [
    {
      to: localePath('/auth/login'),
      icon: 'i-charm-key',
      label: i18n.t('nav_login'),
    },
    {
      to: localePath('/auth/register'),
      icon: 'i-charm-notes',
      label: i18n.t('nav_register'),
    },
  ]
  const staticItems: Array<NavigationMenuItem> = [
    {
      to: localePath('/'),
      icon: 'i-charm-home',
      label: i18n.t('nav_home'),
    },
    {
      to: localePath('/explore'),
      icon: 'i-charm-map',
      label: i18n.t('nav_explore'),
    },
    {
      to: localePath('/rules'),
      icon: 'i-charm-book',
      label: i18n.t('nav_rules'),
    },
  ]
  const helpItem: NavigationMenuItem = {
    label: i18n.t('nav_help'),
    icon: 'i-charm-help',
    disabled: true
  }
  items.push([], staticItems)
  items.push(userStore.user ? [userItem] : authItems)
  items.push([i18nItem, helpItem])
  return items
})

const active = ref()
const collapse = ref(false)

defineShortcuts({
  1: () => active.value = '0',
  2: () => active.value = '1',
  3: () => active.value = '2',
})

</script>

<template>
  <div>
    <nav class="flex flex-col gap-[4vh] mt-4">
        <!--UButton
            :icon="'i-charm-arrow-' + (collapse ? 'right' : 'left' )"
            class="ms-auto me-4"
            size="md"
            color="primary"
            variant="link"
            @click="collapse = !collapse"
        /-->
        <h1 class="text-xl text-center font-bold mt-10 mx-4">Bitacora Forum</h1>
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
