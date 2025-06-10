<script setup lang="ts">

import WorldmapSvg from '@/assets/svg/worldmap.svg'
import CountryCard from './CountryCard.vue'

const emit = defineEmits(['clickCountry', 'mouseoverCountry'])

const handleWorldmapSvg = (event: MouseEvent) => {
  const target = event.target as SVGElement
  const cc = target.getAttribute('cc') || target.parentElement?.getAttribute('cc')
  return cc
}

const handleClick = (event: MouseEvent) => {
  const cc = handleWorldmapSvg(event)
  if (cc) emit('clickCountry', event, cc)
}
const handleMouseover = (event: MouseEvent) => {
  const cc = handleWorldmapSvg(event)
  emit('mouseoverCountry', event, cc)
}

</script>

<style scoped>
.worldmap-svg {
  width: 100%;
  height: auto;
  fill: none;
  object-fit: contain;
}
</style>

<template>

<WorldmapSvg class="worldmap-svg" @click="handleClick" @mouseover="handleMouseover" />

</template>
