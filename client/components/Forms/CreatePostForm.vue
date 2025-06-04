<script setup lang="ts">
import * as z from 'zod'
import { UButton, UFormField, UInput } from '#components';
import CountryChooser from './Fields/CountryChooser.vue';
import { PostModel, type NewPostModel } from '@/models/post'

const modelValue = defineModel<boolean>()

const emit = defineEmits<{
  (e: 'close'): void
}>()

const userStore = useUserStore()


const schema = z.object({
  title: z.string().min(8, 'Must be at least 8 characters'),
  description: z.string().optional(),
  content: z.string().min(25, 'Must be at least 25 characters'),
  country: z.string().nonempty('Se debe relacionar con un país'),
})

type Schema = z.output<typeof schema>

const props = withDefaults(defineProps<{ initialData?: any }>(), { initialData: {} })

const state = reactive<Partial<Schema>>({ ...props.initialData })

async function onSubmit(resultState: Partial<Schema>) {
  const parsed = schema.parse(resultState)
  if (parsed && userStore.user) {
    const res = await PostModel.createPost({
      title: parsed.title,
      description: parsed.description,
      content: parsed.content,
      country_alpha: parsed.country,
      owner_id: userStore.user.id,
    })
    if (res.success) {
      refreshNuxtData()
      modelValue.value = false
    }
  }
}

watch(modelValue, () => emit('close'))

</script>

<template>

<UForm :schema="schema" :state="state" class=" flex flex-col space-y-4">

    <USlideover v-model:open="modelValue" :dismissible="false" :modal="false" :ui="{ content: 'w-[88%] max-w-none', }">
        <template #title>
            <h1 class="text-2xl font-bold">Create new discussion</h1>
        </template>
        <template #footer>
            <UButton color="primary" type="submit" :disabled="!schema.safeParse(state).success" @click="onSubmit(state)">Publicar</UButton>
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
