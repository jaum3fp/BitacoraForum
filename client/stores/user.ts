import { jwtDecode } from "jwt-decode";
import { defineStore } from "pinia";
import { UserModel, type UserModelType } from "~/models/user";


interface UserState {
  user: UserModelType | null
}

export const useUserStore = defineStore('user', {
  state: (): UserState => ({
    user: null
  }),
  getters: {
    loadAndGet: (state: UserState) => {
      return async () => {
        state.user = await UserModel.getOwnData()
        return state.user
      }
    }
  },
  actions: {
    async load() {
      this.user = await UserModel.getOwnData()
    }
  },
  persist: true,
})
