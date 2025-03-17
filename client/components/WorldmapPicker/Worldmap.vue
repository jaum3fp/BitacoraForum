<script setup lang="ts">

import WorldmapSvg from '@/assets/svg/worldmap.svg'
import CountryCard from './CountryCard.vue'

const emit = defineEmits(['selectedCountry'])

function onClickWorldmapSVG(event: Event) {
  const target = event.target as SVGElement
  const cc = target.getAttribute('cc') || target.parentElement?.getAttribute('cc')
  if (cc) {
    console.log('Emit selectedCountry', cc)
    emit('selectedCountry', cc)
  }
}

</script>

<style scoped>
.worldmap-svg {
  width: 100%;
  height: auto;
  fill: none;
}
</style>

<template>

  <WorldmapSvg class="worldmap-svg" @click="onClickWorldmapSVG" />
   
</template>


