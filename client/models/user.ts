import { API } from "~/consts"
import { useUserStore } from "../stores/user";

type UserModelType = {
  id: number;
  username: string;
  email: string;
  name: string | null;
  surnames: string | null;
  profile_img: string | null;
}

type UserRegisterData = UserModelType & {
  password: string;
  repeatPassword: string;
}


const UserModel = {

  register: async (formData: UserRegisterData): Promise<boolean> => {
    try {
      const response = await useApiCall("bitacoraForum", "auth/register", {
        method: "POST",
        body: formData,
      })

      if (response.success) {
        useUserStore().user = await UserModel.getOwnData()
        if (!useUserStore().user) {
          return false
        }
      }

      return response.success
    } catch (error) {
      console.error("No se ha podido registrar el usuario")
      return false
    }
  },

  login: async (formData: any): Promise<boolean> => {
    try {
      const response = await useApiCall("bitacoraForum", "auth/login", {
        method: "POST",
        body: formData,
      })

      if (response.success) {
        useUserStore().user = await UserModel.getOwnData()
        if (!useUserStore().user) {
          return false
        }
      }

      return response.success
    } catch (error) {
      console.error("No se ha podido iniciar sesión")
      return false
    }
  },

  logout: async (): Promise<boolean> => {
    try {
      const response = await useApiCall("bitacoraForum", "auth/logout", {
        method: "POST",
      })

      if (response.success) {
        useUserStore().$reset()
        if (useUserStore().user) {
          return false
        }
        refreshNuxtData()
      }

      return response.success
    } catch (error) {
      console.error("No se ha podido cerrar sesión")
      return false
    }
  },

  getUserData: async (username: string): Promise<UserModelType | null> => {
    try {
      const response = await useApiCall("bitacoraForum", ("user/" + username), {
        method: "GET",
      })
      console.info(response)
      return response
    } catch (error) {
      console.error("No se ha podido obtener el usuario")
      return null
    }
  },

  getOwnData: async (): Promise<UserModelType | null> => {
    try {
      const response = await useApiCall("bitacoraForum", "user/me", {
        method: "GET",
      })
      console.info(response)
      return response
    } catch (error) {
      console.error("No se ha podido obtener el usuario")
      return null
    }
  },

  updateUserData: async (username: string, data: any): Promise<any | null> => {
    try {
      const response = await useApiCall("bitacoraForum", ("user/me"), {
        method: "PUT",
        body: data,
      })
      console.info(response)
      return response
    } catch (error) {
      console.error("No se ha podido actualizar el usuario")
      return null
    }
  },

} as const

type IUserModel = typeof UserModel


export { UserModel, type IUserModel, type UserModelType }
