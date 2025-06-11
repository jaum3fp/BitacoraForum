<script setup lang="ts">
import * as z from 'zod'
import type { FormSubmitEvent } from '@nuxt/ui'
import PasswordField from './Fields/PasswordField.vue'
import ProfileImageField from './Fields/ProfileImageField.vue'
import { regexes } from '~/consts/regex';
import StrengthIndicator from './Fields/StrengthIndicator.vue';
import { UserModel } from '~/models/user';

const modelValue = defineModel<boolean>()

const emit = defineEmits<{
  (e: 'close'): void
}>()

const userStore = useUserStore()
const i18n = useI18n()

const zodPasswordString = z.string()
  .regex(regexes.password.medium, 'Password must contain at least one uppercase letter, one lowercase letter, one number, and one special character')

const schema = z.object({
  profile_img: z.string().nullable(),
  username: z.string().min(8, 'Must be at least 8 characters'),
  password: zodPasswordString,
  repeatPassword: zodPasswordString,
  email: z.string().email("El formato de este email es incorrecto."),
  name: z.string().nullable(),
  surnames: z.string().nullable(),
})

type Schema = z.output<typeof schema>

const state = reactive<Partial<Schema>>({ ...userStore.user })
const initialState = reactive<Partial<Schema>>({ ...userStore.user })

const alteredState = computed(() => JSON.stringify(state) != JSON.stringify(initialState))

async function onSubmit(resultState: Partial<Schema>) {
  if (userStore.user) {
    const success = await UserModel.updateUserData(
      userStore.user.username,
      resultState,
    )
    modelValue.value = false
  }
}

watch(modelValue, () => emit('close'))

</script>

<template>

<UForm :schema="schema" :state="state" class=" flex flex-col space-y-4">

    <USlideover v-model:open="modelValue">
        <template #body>
            <div class="flex flex-col gap-6">

                <div class="flex gap-4">
                    <UFormField name="profile_img">
                        <ProfileImageField v-model="state.profile_img" />
                    </UFormField>
                    <UFormField :label="i18n.t('form_username')" name="username" class="w-full">
                        <UInput disabled v-model="state.username" class="w-full" />
                    </UFormField>
                </div>

                <UFormField :label="i18n.t('form_email')" name="email">
                    <UInput v-model="state.email" class="w-full" />
                </UFormField>

                <UFormField :label="i18n.t('form_name')" name="name">
                    <UInput v-model="state.name" class="w-full" />
                </UFormField>

                <UFormField :label="i18n.t('form_surnames')" name="surnames">
                    <UInput v-model="state.surnames" class="w-full" />
                </UFormField>

                <UButton class="me-auto">{{ i18n.t('user_profile_form_change_password') }}</UButton>

            </div>
        </template>
        <template v-show="alteredState" #footer>
            <UButton color="primary" @click="onSubmit(state)">{{ i18n.t('user_profile_form_submit_button') }}</UButton>
        </template>
    </USlideover>

</UForm>

</template>
