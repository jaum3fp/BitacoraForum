<script setup lang="ts">
import { computed, reactive, watch } from 'vue'

interface CountryCardProps {
    country: string
}

interface CountryData {
    id: number
    name: string
    flag: string
    total_posts: number
}

const props = defineProps<CountryCardProps>()

const { data: countryData} = await useAsyncData<CountryData>(
    'countryData',
    async () => {
        const res = await fetch(`http://localhost:8080/api/v1/country/${props.country.toUpperCase()}`)
        if (!res.ok) {
            console.error("No se ha podido obtener la información del país")
            return
        }
        return await res.json()
    },
    {
        watch: [() => props.country],
        immediate: true
    }
)

const flagapiUrl = computed(() => `https://flagsapi.com/${props.country.toUpperCase()}/flat/64.png`)

</script>

<template>
    <div class="flex items-center p-4 bg-slate-500 shadow-md rounded-lg">
        <img :src="flagapiUrl" alt="flag" />
        <div>
            <h1>{{ countryData?.name }}</h1>
            <p>Nº Posts: {{ countryData?.total_posts }}</p>
        </div>
    </div>
</template>

<style scoped>

</style>