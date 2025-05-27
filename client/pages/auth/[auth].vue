<script setup lang="ts">
import LoginForm from '~/components/Forms/LoginForm.vue'
import RegisterForm from '~/components/Forms/RegisterForm.vue'


const param = useRoute().params.auth as string || 'login'

definePageMeta({
    layout: 'center',
    validate: async (route) => {
      const validParams = ['login', 'register']
      return validParams.includes(route.params.auth as string)
    }
})

const forms = {
  login: LoginForm,
  register: RegisterForm,
}

const formComponent = forms[param as keyof typeof forms] || LoginForm


</script>

<template>
    <component v-if="param === 'login' || param === 'register'" :is="formComponent" />
</template>
