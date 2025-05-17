import { useApiCall } from "~/composables/useApiCall"
import { API } from "~/consts"


const PostModel = {

    getAllPosts: async (country: string | null): Promise<any> => {
      try {
        //const posts = await $fetch(API.bitacoraForum.posts.all + (country ? '/' + country : ''))
        const posts = await useApiCall("bitacoraForum", "post/all" + (country ? '/' + country : ''))

        return posts || []
      } catch (error) {
        console.error("No se han podido obtener todos los posts:", error)
        return []
      }
    }
} as const

type IPostModel = typeof PostModel


export { PostModel, type IPostModel }
