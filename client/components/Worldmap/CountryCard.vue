<script setup lang="ts">
import { computed, reactive, watch } from 'vue'
import { CountryModel, type CountryModelType } from '@/models/country'

interface CountryCardProps {
    cc: string
    name: string
    flag: string
}

const props = defineProps<CountryCardProps>()

const { data: postsNumber } = useAsyncData(
  'postsNumber',
  async () => await useApiCall("bitacoraForum", "post/count/" + props.cc),
  { server: false }
)

</script>

<template>
    <div v-if="props.cc" class="flex items-center p-4 bg-slate-500 shadow-md rounded-lg">
        <img :src="props.flag" alt="flag" class="w-20" />
        <div>
            <h1>{{ props.name }}</h1>
            <p>Nº Posts: {{ postsNumber }}</p>
        </div>
    </div>
    <USkeleton v-else />
</template>

<style scoped>

</style>
