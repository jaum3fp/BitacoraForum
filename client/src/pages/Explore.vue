<script setup lang="ts">
import Worldmap from '../components/WorldmapPicker/Worldmap.vue';
import CountryCard from '../components/WorldmapPicker/CountryCard.vue';
import { ref } from 'vue';

const country = ref('es')

function onSelectedCountry(cc: string) {
    console.log('Selected country:', cc)
    country.value = cc
}
</script>


<template>

<div>
    <CountryCard :country="country" />
    <Worldmap @selected-country="onSelectedCountry" />
</div>

</template>


<style scoped>
</style>
