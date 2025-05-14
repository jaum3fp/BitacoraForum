import { jwtDecode } from "jwt-decode";
import { defineStore } from "pinia";
import type { UserModelType } from "~/models/user";


export const useUserStore = defineStore('user', () => {
  const accessCookie = useCookie('access-token', {
    secure: true,
    //httpOnly: true,
    maxAge: 60 * 60 * 24 * 7,
    sameSite: "strict",
    path: "/"
  })

  let successLogin = ref(false)

  const user: UserModelType = {
    username: "",
    email: "",
    name: null,
    surnames: null
  }

  function reset() {
    successLogin.value = false
    accessCookie.value = null
    user.username = "",
    user.email = "",
    user.name = null,
    user.surnames = null
  }

  function setJWTData(token: string) {
    accessCookie.value = token

    try {
      const decodedJWT = jwtDecode<UserModelType>(token)
      console.info(decodedJWT)
      user.username = decodedJWT.username ?? ''
      user.email = decodedJWT.email ?? ''
      user.name = decodedJWT.name ?? null
      user.surnames = decodedJWT.surnames ?? null
      if (user.username === '') throw Error("No se han podido obtener datos cruciales")
      successLogin.value = true
    } catch (error) {
      console.error('Error decoding user data:', error)
      successLogin.value = false
    }
  }

  return { accessCookie, user, successLogin, setJWTData, reset }
})
