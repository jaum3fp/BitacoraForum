<script setup lang="ts">
import { USelect } from '#components';

const page = defineModel<number>('page')
const itemsPerPage = defineModel<number>('itemsPerPage')
const options = [5, 10, 15, 20, 25, 50, 100]

</script>

<template>

<div class="paginator flex gap-2">
    <UPagination v-model:page="page" :items-per-page="itemsPerPage" v-bind="$attrs" />
    <USelect v-model="itemsPerPage" :items="options" class="w-18" />
</div>

</template>
