import { API } from "~/consts"
import { useUserStore } from "../stores/user";

type UserModelType = {
  username: string;
  email: string;
  name: string | null;
  surnames: string | null;
}

type UserRegisterData = UserModelType & {
  password: string;
  repeatPassword: string;
}

const UserModel = {

  register: async (formData: UserRegisterData): Promise<boolean> => {
    try {
      const response: { token: string } = await $fetch(API.bitacoraForum.auth.register, {
        method: "POST",
        body: formData,
      })
      useUserStore().setJWTData(response.token)
      return useUserStore().successLogin
    } catch (error) {
      console.error("No se ha podido registrar el usuario")
      return false
    }
  },

  login: async (formData: any): Promise<boolean> => {
    try {
      const response: { token: string } = await $fetch(API.bitacoraForum.auth.login, {
        method: "POST",
        body: formData,
      })
      useUserStore().setJWTData(response.token)
      return useUserStore().successLogin
    } catch (error) {
      console.error("No se ha podido iniciar sesión")
      return false
    }
  },

  logout: async () => {
    console.info('logout')
    useUserStore().reset()
  }

} as const

type IUserModel = typeof UserModel


export { UserModel, type IUserModel, type UserModelType }
