export default defineNuxtRouteMiddleware((to) => {
  const cookie = useCookie("access-token")

  if (cookie.value) {
    useUserStore().user = {
      email: '',
      name: '',
      surnames: '',
      username: ''
    }
  }

})
