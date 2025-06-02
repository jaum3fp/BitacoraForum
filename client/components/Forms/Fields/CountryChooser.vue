<script setup lang="ts">
import { UAvatar, USelectMenu } from '#components';

const countriesStore = useCountriesStore()
const modelValue = defineModel<string>()

const countries = Object.entries(countriesStore.countries)

const items = countries.map(([key, val]) => ({
  label: val.name,
  flag: val.flagSvg,
  id: key,
}))

const getFlag = (cc: string) => {
  const t = countries.find(([key]) => key === cc )
  return t?.[1].flagSvg
}

</script>

<template>
    <div class="countrychooser">
        <USelectMenu v-model="modelValue" :items="items" value-key="id" class="w-48">
            <template #leading="{ modelValue, ui }">
                <UAvatar v-if="modelValue" :src="getFlag(modelValue)" size="sm" :class="ui.leadingIcon()" />
                <UIcon v-else name="i-charm-globe" :class="ui.leadingIcon()" />
            </template>
            <template #item-leading="{ item }">
                <UAvatar :src="item.flag" size="sm" />
            </template>
        </USelectMenu>
    </div>
</template>
