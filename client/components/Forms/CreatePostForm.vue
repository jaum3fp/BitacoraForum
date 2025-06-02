<script setup lang="ts">
import * as z from 'zod'
import type { FormSubmitEvent } from '@nuxt/ui'
import PasswordField from './Fields/PasswordField.vue'
import ProfileImageField from './Fields/ProfileImageField.vue'
import { regexes } from '~/consts/regex';
import StrengthIndicator from './Fields/StrengthIndicator.vue';
import { UserModel } from '~/models/user';
import { UButton, UFormField, UInput } from '#components';
import CountryChooser from './Fields/CountryChooser.vue';

const modelValue = defineModel<boolean>()

const emit = defineEmits<{
  (e: 'close'): void
}>()

const userStore = useUserStore()


const schema = z.object({
  title: z.string().min(8, 'Must be at least 8 characters'),
  description: z.string().nullable().optional(),
  content: z.string().min(25, 'Must be at least 25 characters'),
  country: z.string(),
})

type Schema = z.output<typeof schema>

const state = reactive<Partial<Schema>>({})

async function onSubmit(resultState: Partial<Schema>) {

}

watch(modelValue, () => emit('close'))

</script>

<template>

<UForm :schema="schema" :state="state" class=" flex flex-col space-y-4" @submit="onSubmit">

    <USlideover v-model:open="modelValue" :dismissible="false" :modal="false" :ui="{ content: 'w-[88%] max-w-none', }">
        <template #title>
            <h1 class="text-2xl font-bold">Create new discussion</h1>
        </template>
        <template #footer>
            <UButton color="primary" :disabled="!schema.safeParse(state).success">Publicar</UButton>
            <UButton color="neutral" variant="outline" @click="modelValue = false">Cancelar</UButton>
        </template>
        <template #body>
            <div class="flex flex-col gap-6">

                <div class="flex gap-6">
                    <UFormField label="Title" name="title" class="">
                        <UInput v-model="state.title" class="w-96" />
                    </UFormField>

                    <UFormField label="Country" name="country">
                        <CountryChooser v-model="state.country" />
                    </UFormField>
                </div>

                <UFormField label="Description" name="description">
                    <UTextarea v-model="state.description" class="w-full" />
                </UFormField>

                <UFormField label="Content" name="content">
                    <UTextarea v-model="state.content" class="w-full" />
                </UFormField>

            </div>
        </template>
    </USlideover>

</UForm>

</template>
