<script setup lang="ts">
import * as z from 'zod'
import type { FormSubmitEvent } from '@nuxt/ui'
import PasswordField from './Fields/PasswordField.vue'
import { regexes } from '~/consts/regex';
import StrengthIndicator from './Fields/StrengthIndicator.vue';
import { UserModel } from '~/models/user';


const zodPasswordString = z.string()
  .regex(regexes.password.medium, 'Password must contain at least one uppercase letter, one lowercase letter, one number, and one special character')

const schema = z.object({
  username: z.string().min(8, 'Must be at least 8 characters'),
  password: zodPasswordString,
  repeatPassword: zodPasswordString,
  email: z.string().email("El formato de este email es incorrecto."),
  name: z.string().nullable(),
  surnames: z.string().nullable(),
})

schema.refine((it) => !('repeatPassword' in it) || it.password === it.repeatPassword, {
  message: "Passwords don't match",
  path: ['repeatPassword'],
})

type Schema = z.output<typeof schema>
const state = reactive<Partial<Schema>>({})

async function onSubmit(event: FormSubmitEvent<Schema>) {
  const success: boolean = await UserModel.register(event.data)
  if (success) {
    await navigateTo('/')
  }
}

</script>

<template>
    <div class="register-form w-[30%] my-auto bg-black opacity-60 p-14 rounded-lg">
        <UForm :schema="schema" :state="state" class=" flex flex-col space-y-4" @submit="onSubmit">

            <UFormField label="Username" name="username">
                <UInput v-model="state.username" class="w-full" />
            </UFormField>

            <UFormField label="Password" name="username">
                <PasswordField v-model="state.password" class="mb-2" />
                <StrengthIndicator v-model="state.password" />
            </UFormField>

            <UFormField label="Repeat" name="repeatPassword">
                <PasswordField v-model="state.repeatPassword" />
            </UFormField>

            <UFormField label="Email" name="email">
                <UInput v-model="state.email" class="w-full" />
            </UFormField>

            <UFormField label="Name" name="name">
                <UInput v-model="state.name" class="w-full" />
            </UFormField>

            <UFormField label="Surnames" name="surnames">
                <UInput v-model="state.surnames" class="w-full" />
            </UFormField>

            <UButton color="primary" type="submit" :disabled="!schema.safeParse(state).success">Register</UButton>

        </UForm>
    </div>
</template>
