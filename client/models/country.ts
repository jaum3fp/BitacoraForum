import type { any, string } from "zod"
import { API } from "~/consts"

interface CountryModelType {
  name: string,
  posts?: number,
  flagSvg: string,
}

const CountryModel = {

  getCountriesData: async (): Promise<Record<string, CountryModelType>> => {
    try {
      const countries = await useApiCall("restcountries", "all?fields=flags,cca2,name")

      return countries.reduce((acc: any, country: any) => {
        acc[country.cca2] = ({ name: country.name.common, flagSvg: country.flags.svg })
        return acc
      }, {})
    } catch (error) {
      console.error("No se ha podido obtener la información del país:", error)
      return {}
    }
  },

  getCountryDataByAlpha: async (alpha: string, query: string = ''): Promise<CountryModelType> => {
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

  getCountryFlags: async (alphas: Array<string>): Promise<object> => {
    try {
      const countries = await useApiCall("restcountries", ("alpha?fields=flags,cca2&codes=" + alphas.join()))
      return countries.reduce((acc: any, it: any) => {
        acc[it.cca2] = it.flags.svg
        return acc
      }, {})
    } catch (error) {
      console.error("No se ha podido obtener las banderas:", error)
      return {}
    }
  }

} as const

type ICountryModel = typeof CountryModel


export { CountryModel, type ICountryModel, type CountryModelType }
