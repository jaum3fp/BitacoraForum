<script setup lang="ts">

interface SearchBarProps {
  searchTimeout?: number
  clearable?: boolean
  filter?: SearchFilter
}

interface SearchFilter {
  component: any
  args?: any
}

const props = withDefaults(defineProps<SearchBarProps>(), {
  searchTimeout: 1000,
  clearable: true,
})

const emit = defineEmits<{
  (e: 'search', value: string): void
}>()

const searchInput = ref("")

const searching = ref(false)
let searchTimeout: NodeJS.Timeout | null = null

watch(searchInput, (newVal) => {
  if (searchTimeout) {
    clearTimeout(searchTimeout)
    searchTimeout = null
  }
  if (!newVal || newVal.trim() === '') {
    searching.value = false
    emit("search", newVal)
    return
  }
  searching.value = true
  searchTimeout = setTimeout(() => {
    searching.value = false
    emit("search", newVal)
    searchTimeout = null
  }, props.searchTimeout)
})

</script>

<template>
    <div class="search-bar">
        <UInput v-model="searchInput"
            size="xl"
            :loading="searching"
            icon="i-charm-search"
            placeholder="Search..."
            class="w-full"
            v-bind="$attrs"
        >
            <!--template v-if="props.clearable || props.filter" #trailing>
                <UButton v-if="props.clearable && searchInput.length > 0"
                    color="neutral"
                    variant="link"
                    size="sm"
                    icon="i-charm-circle-cross"
                    aria-label="Clear input"
                    @click="searchInput = ''"
                />
                <component v-if="props.filter" :is="props.filter.component" v-bind="props.filter.args" />
            </template-->
        </UInput>
    </div>
</template>
