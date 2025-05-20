<script setup lang="ts">
import Worldmap from '../components/Worldmap/Worldmap.vue';
import CountryCard from '../components/Worldmap/CountryCard.vue';
import { ref } from 'vue';


const cardCountry = ref<string | null>(null)
const showCountryPosts = reactive<{ country: string, show: boolean }>({
  country: '',
  show: false,
})

defineShortcuts({
  o: () => showCountryPosts.show = !showCountryPosts.show
})

const mapCoords = ref({ x: -1, y: -1 })

function onClickCountry(event: MouseEvent, cc: string) {
  cardCountry.value = null
  showCountryPosts.country = cc
  showCountryPosts.show = true
}

function onMouseoverCountry(event: MouseEvent, cc: string) {
    const { offsetX: x, offsetY: y } = event
    mapCoords.value = { x, y }
    cardCountry.value = cc
}

</script>


<template>

<div class="explorepage relative w-full h-screen">

    <Worldmap class="absolute inset-o w-full h-full"
        @click-country="onClickCountry"
        @mouseover-country="onMouseoverCountry"
    />

    <CountryCard v-if="cardCountry"
        :country="cardCountry"
        :style="{ top: `${mapCoords.y}px`, left: `${mapCoords.x}px` }"
        class="absolute"
    />

    <USlideover v-model:open="showCountryPosts.show"
          side="right"
          :overlay="true"
          :ui="{ content: 'w-[60%] max-w-none', }"
        >
        <template #body>
            <PostsList :country="showCountryPosts.country" />
        </template>
    </USlideover>
</div>

</template>
