<script setup lang="ts">
import type { BreadcrumbItem } from '@nuxt/ui'

const route = useRoute()
const pathNames = computed(() => route.path.split('/'))
const localePath = useLocalePath()
const i18n = useI18n()

const items = computed<BreadcrumbItem[]>(() => pathNames.value
  .filter((it, idx) =>
    !(it === '' && idx === 1) &&
    !i18n.locales.value.map(i => i.code).includes(it as any)
  )
  .map((it, idx) => {
    if (it === '') {
      return ({
        label: i18n.t('nav_home'),
        icon: 'i-charm-home',
        to: localePath('/')
      })
    } else if (it === 'explore') {
      return ({
        label: i18n.t('nav_explore'),
        icon: 'i-charm-map',
        to: localePath('/explore')
      })
    } else if (it === 'rules') {
      return ({
        label:  i18n.t('nav_rules'),
        icon: 'i-charm-book',
        to: localePath('/rules')
      })
    } else if (it === 'discussions') {
      return ({
        label: i18n.t('nav_discussions'),
        to: localePath('/')
      })
    } else if (it === 'user') {
      return ({
        label: i18n.t('nav_user'),
        to: localePath('/')
      })
    } else {
      return ({
        label: i18n.t('nav_' + it) !== 'nav_' + it ? i18n.t('nav_' + it) : it,
        to: localePath(pathNames.value.length === (idx + 1) ? route.path : '/')
      })
    }
  }
))

</script>

<template>
  <UBreadcrumb :items="items" />
</template>
