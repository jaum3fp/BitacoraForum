import { UserModel } from "~/models/user"

export default defineNuxtRouteMiddleware(async (to) => {
  const cookie = useCookie("access-token")
  const userCookie = useCookie("user")

  if (cookie.value && !userCookie.value) {
    useUserStore().user = await UserModel.getOwnData()
  }

})
