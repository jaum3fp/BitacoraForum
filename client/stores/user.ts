import { jwtDecode } from "jwt-decode";
import { defineStore } from "pinia";
import type { UserModelType } from "~/models/user";


interface UserState {
  user: UserModelType | null
}

export const useUserStore = defineStore('user', {
  state: (): UserState => ({
    user: null
  }),
  persist: true,
})
