<script setup lang="ts">
import Worldmap from '../components/Worldmap/Worldmap.vue';
import CountryCard from '../components/Worldmap/CountryCard.vue';
import { ref } from 'vue';
import { CountryModel } from '~/models/country';
import type { CountryModelType } from '~/models/country';


const countriesStore = useCountriesStore()
const cardCountry = ref<string | null>(null)
const showCountryPosts = reactive<{ country: string, show: boolean }>({
  country: '',
  show: false,
})
const mapCoords = ref({ x: -1, y: -1 })

defineShortcuts({
  o: () => showCountryPosts.show = !showCountryPosts.show
})

const currentCountry = computed(() => {
  const cc = cardCountry.value
  if (!cc) return null
  return countriesStore.getCountryByAlpha(cc.toUpperCase())
})

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

useAsyncData(
  'countriesData',
  async () => countriesStore.countries = await CountryModel.getCountriesData(),
  { server: false }
)

</script>


<template>

<div class="explorepage relative w-full h-screen">

    <Worldmap class="absolute inset-o w-full h-full"
        @click-country="onClickCountry"
        @mouseover-country="onMouseoverCountry"
    />

    <CountryCard v-if="currentCountry && cardCountry"
        :cc="cardCountry"
        :name="currentCountry.name"
        :flag="currentCountry.flagSvg"
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
