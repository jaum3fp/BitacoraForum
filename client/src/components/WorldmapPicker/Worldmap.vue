<script setup lang="ts">

import { onMounted, useTemplateRef } from 'vue'

const worldmapRef = useTemplateRef('worldmap')
const emit = defineEmits(['selectedCountry'])

onMounted(() => {
  worldmapRef.value?.addEventListener('load', function() {
    this.getSVGDocument()?.addEventListener('click', e => {
      const target = e.target as SVGElement
      const cc = target.getAttribute('cc') || target.parentElement?.getAttribute('cc')
      if (cc) {
        console.log('Emit selectedCountry', cc)
        emit('selectedCountry', cc)
      }
    })
  })
})

</script>

<style scoped>
embed {
  width: 100%;
  height: auto;
  background-color: rgba(0, 0, 0, 0);
}
</style>

<template>

<div class="grid-container">
  <embed ref="worldmap" src="/src/assets/worldmap.svg" type="image/svg+xml" />
</div>

</template>


