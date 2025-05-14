import { API } from "~/consts"

interface CountryData {
  name: string,
  posts: number,
  flagSvg: string,
}

const CountryModel = {

  // TODO: TIPAR CON ISO 3166-1 alpha-2
  getCountryDataByAlpha: async (alpha: string, query: string = ''): Promise<CountryData | undefined> => {
    try {
      const country = await $fetch<any>(`${API.restcountries.alpha}/${alpha}?${query}`)
      const countyPosts = await $fetch<number>(`${API.bitacoraForum.posts.countCountry}${alpha}`)

      return {
        name: country.name.common,
        flagSvg: country.flags.svg,
        posts: countyPosts || 0,
      }
    } catch (error) {
      console.error("No se ha podido obtener la información del país:", error)
      return undefined
    }
  }
} as const

type ICountryModel = typeof CountryModel


export { CountryModel, type ICountryModel, type CountryData }
