<script setup lang="ts">
import Worldmap from '../components/WorldmapPicker/Worldmap.vue';
import CountryCard from '../components/WorldmapPicker/CountryCard.vue';
import { ref } from 'vue';

const country = ref('es')

function onSelectedCountry(cc: string) {
    country.value = cc
}
</script>


<template>

    <div>
        <CountryCard :country="country" />
        <div class="m-5">
            <Worldmap @selected-country="onSelectedCountry" />
        </div>
    </div>

</template>


<style scoped>
</style>
