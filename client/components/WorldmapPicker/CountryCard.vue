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

const { country } = defineProps<CountryCardProps>()
const countryData = reactive<CountryData>({
    id: 0,
    name: '',
    flag: '',
    total_posts: 0
})
const flagapiUrl = computed(() => `none`) // https://flagsapi.com/${country.toUpperCase()}/flat/64.png

async function fetchCountryPosts() {
    const res = await fetch(`http://localhost:8080/api/v1/country/${country.toUpperCase()}`)
    const resJSON: CountryData = await res.json()
    countryData.name = resJSON.name
    countryData.total_posts = resJSON.total_posts
}

async function useFlagsapi(cc: string) {
    const res = await fetch(`https://restcountries.com/v3.1/alpha/${cc.toUpperCase()}`)
    const resJSON = await res.json()
    return resJSON.flags[0]

}

watch(() => country, fetchCountryPosts, { immediate: true })

</script>

<template>
    <div>
        <div>
            
        </div>
        <div>
            <h1>{{ countryData.name }}</h1>
            <p>Nº Posts: {{ countryData.total_posts }}</p>
        </div>
    </div>
</template>

<style scoped>

</style>