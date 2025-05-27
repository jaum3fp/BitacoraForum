export default defineNuxtPlugin(() => {
  const userStore = useUserStore()

  userStore.$subscribe((_, state) => {
    if (!state.user) refreshNuxtData()
  })
})
