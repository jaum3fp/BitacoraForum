import { API } from "~/consts"


const PostModel = {

  getAllPosts: async () => {
    try {
      const res = await fetch(API.bitacoraForum.posts.all)
      return await res.json()
    } catch (error) {
      console.error("No se han podido obtener todos los posts")
    }
  }

} as const

type IPostModel = typeof PostModel


export { PostModel, type IPostModel }
