<script setup lang="ts">

import WorldmapSvg from '@/assets/svg/worldmap.svg'
import CountryCard from './CountryCard.vue'

const emit = defineEmits(['clickCountry'])

function handleWorldmapSVG(event: MouseEvent) {
  const target = event.target as SVGElement
  const cc = target.getAttribute('cc') || target.parentElement?.getAttribute('cc')

  console.log('Emit selectedCountry', cc)
  emit('clickCountry', event, cc)
  
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

  <WorldmapSvg class="worldmap-svg" @mouseover="handleWorldmapSVG" />
   
</template>


