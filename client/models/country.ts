import { API } from "~/consts"

interface CountryData {
  name: string,
  posts: number,
  flagSvg: string,
}

const CountryModel = {

  // TODO: TIPAR CON ISO 3166-1 alpha-2
  getCountryDataByAlpha: async (alpha: string, query: string = ''): Promise<CountryData> => {
    try {
      const country = await useApiCall("restcountries", ("alpha/" + alpha + '?' + query))
      const countryPosts = await useApiCall("bitacoraForum", "post/count/" + alpha)

      return {
        name: country.name.common,
        flagSvg: country.flags.svg,
        posts: countryPosts || 0,
      }
    } catch (error) {
      console.error("No se ha podido obtener la información del país:", error)
      return {
        name: "",
        posts: 0,
        flagSvg: "",
      }
    }
  },

  getCountryFlag: async (alpha: string): Promise<{ flagSvg: string }> => {
    try {
      const country = await useApiCall("restcountries", ("alpha/" + alpha + "?fields=flags"))

      return {
        flagSvg: country.flags.svg,
      }
    } catch (error) {
      console.error("No se ha podido obtener la información del país:", error)
      return { flagSvg: "" }
    }
  }

} as const

type ICountryModel = typeof CountryModel


export { CountryModel, type ICountryModel, type CountryData }
