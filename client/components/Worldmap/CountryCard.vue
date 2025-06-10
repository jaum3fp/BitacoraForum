<script setup lang="ts">

interface CountryCardProps {
    cc: string
    name: string
    flag: string
}

const props = defineProps<CountryCardProps>()
const cc = ref(props.cc)

const { data: postsNumber } = useAsyncData(
  'postsNumber',
  async () => await useApiCall("bitacoraForum", "post/count/" + cc.value),
  { server: false }
)

</script>

<template>
    <div v-if="cc !== ''" class="flex items-center text-black gap-3 p-4 bg-primary-500/90 dark:bg-slate-500/70 shadow-md" @click="cc = ''">
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
