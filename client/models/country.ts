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
      const country = await fetch(API.restcountries.alpha + `/${alpha}?${query}`)
      const countyPosts = await fetch(API.bitacoraForum.posts.countCountry + alpha)
      const countryJSON = await country.json()
      const countyPostsJSON = await countyPosts.json()
      return {
        name: countryJSON.name.common,
        flagSvg: countryJSON.flags.svg,
        posts: countyPostsJSON,
      }
    } catch (error) {
      console.error("No se ha podido obtener la información del país")
    }
  }

} as const

type ICountryModel = typeof CountryModel


export { CountryModel, type ICountryModel, type CountryData }
