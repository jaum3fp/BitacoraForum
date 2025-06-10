<script setup lang="ts">
import type { BreadcrumbItem } from '@nuxt/ui'

const route = useRoute()
const pathNames = computed(() => route.path.split('/'))
const localePath = useLocalePath()

const items = computed<BreadcrumbItem[]>(() => pathNames.value
  .filter((it, idx) => !(it === '' && idx === 1))
  .map((it, idx) => {
    if (it === '') {
      return ({
        label: 'Inicio',
        icon: 'i-charm-home',
        to: localePath('/')
      })
    } else if (it === 'explore') {
      return ({
        label: 'Explore',
        icon: 'i-charm-map',
        to: localePath('/explore')
      })
    } else if (it === 'rules') {
      return ({
        label: 'Rules',
        icon: 'i-charm-book',
        to: localePath('/rules')
      })
    } else if (it === 'discussions') {
      return ({
        label: it,
        to: localePath('/')
      })
    } else if (it === 'user') {
      return ({
        label: it,
        to: localePath('/')
      })
    } else {
      return ({
        label: it,
        to: localePath(pathNames.value.length === (idx + 1) ? route.path : '/')
      })
    }
  }
))

</script>

<template>
  <UBreadcrumb :items="items" />
</template>
