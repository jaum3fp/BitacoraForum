import { defineStore } from "pinia";
import { CountryModel, type CountryModelType } from "~/models/country";


interface CountriesState {
  countries: Record<string, CountryModelType>
}

export const useCountriesStore = defineStore('countries', {
  state: (): CountriesState => ({
    countries: {}
  }),
  getters: {
    getCountryByAlpha: (state: CountriesState) => {
      return (alpha: string) => state.countries[alpha]
    },
    getAlphaByName: (state: CountriesState) => {
      return (name: string) => Object.keys(state.countries).find(key => state.countries[key].name.toLowerCase() === name.toLowerCase())
    }
  },
  actions: {
    async load() {
      this.countries = await CountryModel.getCountriesData()
    }
  }
})
