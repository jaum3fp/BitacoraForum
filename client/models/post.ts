import { API } from "~/consts"


const PostModel = {

    getAllPosts: async (): Promise<any | undefined> => {
      try {
        const posts = await $fetch(API.bitacoraForum.posts.all, {
          headers: { 'Authorization': 'Bearer' + useUserStore().accessCookie }
        })

        return posts || []
      } catch (error) {
        console.error("No se han podido obtener todos los posts:", error)
        return undefined
      }
    }
} as const

type IPostModel = typeof PostModel


export { PostModel, type IPostModel }
