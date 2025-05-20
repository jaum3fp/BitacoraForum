import { API } from "~/consts"


export async function useApiCall(
  target: "restcountries" | "bitacoraForum",
  url: string,
  opts: any = {})
{
  const { ssrContext } = useNuxtApp()
  try {

    const defaultOptions: any = {
      baseURL: API[target],
    }

    if (target === "bitacoraForum") {
      defaultOptions.credentials = 'include'
      if (ssrContext) {
        const headers = useRequestHeaders(['cookie'])
        if (headers.cookie) defaultOptions.headers = { cookie: headers.cookie }
      }
    }

    const data = await $fetch<any>(url, { ...defaultOptions, ...opts })
    return data
  } catch (error) {
    console.error("Error obteniendo los datos:", error)
  }
}
