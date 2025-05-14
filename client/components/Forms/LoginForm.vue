<script setup lang="ts">
import * as z from 'zod'
import type { FormSubmitEvent } from '@nuxt/ui'
import { reactive } from 'vue';
import PasswordField from './Fields/PasswordField.vue'
import { regexes } from '../../consts/regex';
import { UserModel } from '../../models/user';


const schema = z.object({
  username: z.string().min(8, 'Must be at least 8 characters'),
  password: z.string()
    .regex(regexes.password.medium, 'Password must contain at least one uppercase letter, one lowercase letter, one number, and one special character')
})

type Schema = z.output<typeof schema>

const state = reactive<Partial<Schema>>({})

async function onSubmit(event: FormSubmitEvent<Schema>) {
  const success: boolean = await UserModel.login(event.data)
  if (success) {
    await navigateTo('/')
  }
}

</script>

<template>
    <div class="login-form my-auto bg-black opacity-60 p-4 rounded-lg">
        <UForm :schema="schema" :state="state" class=" flex flex-col space-y-4" @submit="onSubmit">

            <UFormField label="Username" name="username">
                <UInput v-model="state.username" />
            </UFormField>

            <PasswordField v-model="state.password" />

            <UButton color="primary" type="submit" :disabled="!schema.safeParse(state).success">Login</UButton>

            <p class="text-red" @click="navigateTo('register')">Registrarse</p>

        </UForm>
    </div>
</template>
