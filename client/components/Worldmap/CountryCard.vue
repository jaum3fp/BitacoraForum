<script setup lang="ts">
import { computed, reactive, watch } from 'vue'
import { CountryModel, type CountryData } from '@/models/country'

interface CountryCardProps {
    country: string
}

const props = defineProps<CountryCardProps>()

const { data: countryData} = await useAsyncData<CountryData | undefined>(
    'countryData',
    async () => await CountryModel.getCountryDataByAlpha(props.country, 'fields=name,flags'),
    { watch: [() => props.country], immediate: true }
)

</script>

<template>
    <div v-if="countryData" class="flex items-center p-4 bg-slate-500 shadow-md rounded-lg">
        <img :src="countryData.flagSvg" alt="flag" class="w-20" />
        <div>
            <h1>{{ countryData.name }}</h1>
            <p>Nº Posts: {{ countryData.posts }}</p>
        </div>
    </div>
</template>

<style scoped>

</style>
