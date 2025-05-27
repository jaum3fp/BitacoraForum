<script setup lang="ts">

const modelValue = defineModel<any[] | undefined>()

const items = ref(['Country', 'Author', 'Tags'])

</script>

<template>
    <div class="search-filter">
        <USelect
            v-model="modelValue"
            :multiple="true"
            :items="items"
            icon="i-charm-filter"
            v-bind="$attrs"
        >
        </USelect>
    </div>
</template>
