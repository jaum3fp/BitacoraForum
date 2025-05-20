import { defineStore } from "pinia";
import type { CountryModelType } from "~/models/country";


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
    }
  }
})
