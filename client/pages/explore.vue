<script setup lang="ts">
import Worldmap from '../components/Worldmap/Worldmap.vue';
import CountryCard from '../components/Worldmap/CountryCard.vue';
import { ref } from 'vue';


const country = ref<string | null>(null)

const mapCoords = ref({ x: -1, y: -1 })

function onSelectedCountry(event: MouseEvent, cc: string) {
    const { offsetX: x, offsetY: y } = event
    mapCoords.value = { x, y }
    country.value = cc
}

</script>


<template>

    <div class="relative w-full h-screen">
        <Worldmap class="absolute inset-o w-full h-full" @click-country="onSelectedCountry" />
        <CountryCard 
            v-if="country"
            :country="country"
            :style="{ top: `${mapCoords.y}px`, left: `${mapCoords.x}px` }"
            class="absolute"
        />
    </div>

</template>


<style scoped>
</style>
