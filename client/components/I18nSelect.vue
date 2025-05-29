<script setup lang="ts">
import { avatar } from '#build/ui'

const { locales, locale } = useI18n()
const switchLocalePath = useSwitchLocalePath()
const countriesStore = useCountriesStore()

const getFlag = (cc: string) => countriesStore.getCountryByAlpha(cc.toUpperCase()).flagSvg

const currentLang = ref(locale.value)

const items = ref(locales.value.map(lang => {
  const cc = lang.language?.split('-')[1]
  return {
    label: lang.name,
    value: lang.code,
    avatar: { src: cc ? getFlag(cc) : '' }
  }
}))

const currentAvatar = computed(() => items.value.find(item => item.value === currentLang.value)?.avatar)

</script>

<template>

<div class="i18nselect">
    <USelect v-model="currentLang"
        :items="items"
        :avatar="currentAvatar"
        :ui="{
          trailingIcon: 'group-data-[state=open]:rotate-180 transition-transform duration-200'
        }"
        class="w-full"
        v-bind="$attrs"
        @update:modelValue="navigateTo(switchLocalePath(currentLang))"
    />
</div>

</template>
