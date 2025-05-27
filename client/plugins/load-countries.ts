export default defineNuxtPlugin(async () => {
  const countriesStore = useCountriesStore()

  if (Object.keys(countriesStore.countries).length === 0) {
    await countriesStore.load()
  }
})
