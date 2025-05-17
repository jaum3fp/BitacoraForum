import { API } from "~/consts"


export async function useApiCall(
  target: "restcountries" | "bitacoraForum",
  url: string,
  opts: any = {})
{
  const { ssrContext } = useNuxtApp()
  const request = ssrContext ?
    API[target].ssr.base + url :
    API[target].base + url
  try {
    const data = await $fetch<any>(request, opts)
    return data
  } catch (error) {
    console.error("Error obteniendo los datos:", error)
  }
}
