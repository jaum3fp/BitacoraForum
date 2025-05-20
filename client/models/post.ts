import { useApiCall } from "~/composables/useApiCall"
import { API } from "~/consts"
import { CountryModel } from "./country"


const PostModel = {

    getAllPosts: async (country: string | null): Promise<any> => {
      try {
        const posts = await useApiCall("bitacoraForum", "post/all" + (country ? '/' + country : ''))
        const alphas: Set<string> = new Set(posts.map((post: any) => post.country_alpha))
        const flags: any = await CountryModel.getCountryFlags([...alphas])
        const res = posts.map((post: any) => ({ ...post, flag: flags[post.country_alpha] }))
        return res || []
      } catch (error) {
        console.error("No se han podido obtener todos los posts:", error)
        return []
      }
    }

} as const

type IPostModel = typeof PostModel


export { PostModel, type IPostModel }
