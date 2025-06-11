<script setup lang="ts">
import * as z from 'zod'
import { UButton, UFormField, UInput } from '#components';
import CountryChooser from './Fields/CountryChooser.vue';

const modelValue = defineModel<boolean>()

const i18n = useI18n()

const emit = defineEmits<{
  (e: 'close'): void
  (e: 'submit', state: Schema): void
}>()

const schema = z.object({
  title: z.string().min(8, 'Must be at least 8 characters'),
  description: z.string().optional(),
  content: z.string().min(25, 'Must be at least 25 characters'),
  country_alpha: z.string().nonempty('Se debe relacionar con un país'),
})

type Schema = z.output<typeof schema>

const props = withDefaults(defineProps<{
  title?: string
  initialData?: any
}>(), {
  initialData: {}
})

const state = reactive<Partial<Schema>>({ ...props.initialData })

const handleSubmit = async (state: Partial<Schema>) => {
  const parsed = schema.parse(state)
  if (parsed) emit('submit', parsed)
}

watch(modelValue, () => emit('close'))

</script>

<template>

<UForm :schema="schema" :state="state" class=" flex flex-col space-y-4">

    <USlideover v-model:open="modelValue" :dismissible="false" :modal="false" :ui="{ content: 'w-[88%] max-w-none', }">
        <template #title>
            <h1 class="text-2xl font-bold">{{ props.title }}</h1>
        </template>
        <template #footer>
            <UButton color="primary" :disabled="!schema.safeParse(state).success" @click="handleSubmit(state)">{{ i18n.t('form_publicate') }}</UButton>
            <UButton color="neutral" variant="outline" @click="modelValue = false">{{ i18n.t('form_cancel') }}</UButton>
        </template>
        <template #body>
            <div class="flex flex-col gap-6">

                <div class="flex gap-6">
                    <UFormField :label="i18n.t('form_title')" name="title" class="">
                        <UInput v-model="state.title" class="w-96" />
                    </UFormField>

                    <UFormField :label="i18n.t('form_country')" name="country">
                        <CountryChooser v-model="state.country_alpha" />
                    </UFormField>
                </div>

                <UFormField :label="i18n.t('form_description')" name="description">
                    <UTextarea v-model="state.description" class="w-full" />
                </UFormField>

                <UFormField :label="i18n.t('form_content')" name="content">
                    <UTextarea v-model="state.content" class="w-full" />
                </UFormField>

            </div>
        </template>
    </USlideover>

</UForm>

</template>
