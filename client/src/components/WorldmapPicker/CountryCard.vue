<script setup lang="ts">
import { computed, reactive, watch } from 'vue'

interface CountryCardProps {
    country: string
}

interface CountryData {
    name: string
    total_posts: number
}

const { country } = defineProps<CountryCardProps>()
const countryData = reactive<CountryData>({ name: '', total_posts: 0 })
const flagapiUrl = computed(() => `none`) // https://flagsapi.com/${country.toUpperCase()}/flat/64.png

async function fetchCountryPosts() {
    const res = await fetch(`http://localhost:8080/api/v1/country/${country.toUpperCase()}`)
    const resJSON: CountryData = await res.json()
    countryData.name = resJSON.name
    countryData.total_posts = resJSON.total_posts
}

watch(() => country, fetchCountryPosts, { immediate: true })

</script>

<template>
    <div>
        <div>
            <img :src="flagapiUrl">
        </div>
        <div>
            <h1>{{ countryData.name }}</h1>
            <p>Nº Posts: {{ countryData.total_posts }}</p>
        </div>
    </div>
</template>

<style scoped>

</style>