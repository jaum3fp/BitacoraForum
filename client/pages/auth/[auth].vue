<script setup lang="ts">
import LoginForm from '~/components/Forms/LoginForm.vue'
import RegisterForm from '~/components/Forms/RegisterForm.vue'

definePageMeta({
    layout: 'center',
})

const forms = {
  login: LoginForm,
  register: RegisterForm,
}

const param = useRoute().params.auth || 'login'
const formComponent = forms[param as keyof typeof forms] || LoginForm


</script>

<template>
    <component v-if="param === 'login' || param === 'register'" :is="formComponent" />
</template>
