import { useApiCall } from "~/composables/useApiCall"
import { API } from "~/consts"
import { CountryModel } from "./country"
import buildQueryParams from "~/utils/queryParamsBuilder"
import { id } from "@nuxt/ui/runtime/locale/index.js"


interface PostsFilterType {
  author?: string,
  cc?: string,
  tags?: Array<string>,
}

interface PostModelType {
  id: number,
  created_at: Date,
  updated_at: Date,
  title: string,
  description: string,
  content: string,
  views: number,
  owner_id: number,
  country_alpha: string,
}

interface NewPostModel {
  title: string,
  description?: string,
  content: string,
  country_alpha: string,
  owner_id: number,
}

const PostModel = {

    getAllPosts: async (filter: PostsFilterType): Promise<any> => {
      try {
        const queryParams = buildQueryParams(filter)
        const posts = await useApiCall("bitacoraForum", "post/all" + queryParams)
        if (posts.length <= 0) return posts
        const alphas: Set<string> = new Set(posts.map((post: any) => post.country_alpha))
        const flags: any = await CountryModel.getCountryFlags([...alphas])
        const res = posts.map((post: any) => ({ ...post, flag: flags[post.country_alpha] }))
        return res || []
      } catch (error) {
        console.error("No se han podido obtener todos los posts:", error)
        return []
      }
    },

    getPostById: async (id: number): Promise<any> => {
      try {
        const posts = await useApiCall("bitacoraForum", ("post/" + id))
        if (posts.length <= 0) return posts
        return posts || []
      } catch (error) {
        console.error("No se han podido obtener todos los posts:", error)
        return []
      }
    },

    incrementView: async (id: number): Promise<any> => {
      try {
        const res = await useApiCall("bitacoraForum", ("post/inc-view/" + id), {
          method: 'PATCH'
        })
        return res
      } catch (error) {
        console.error("No se han podido obtener todos los posts:", error)
        return null
      }
    },

    createPost: async (newPost: NewPostModel): Promise<any> => {
      try {
        const res = await useApiCall("bitacoraForum", "/post/", {
          method: 'POST',
          body: newPost
        })
        return res
      } catch (error) {
        console.error("No se han podido obtener todos los posts:", error)
        return null
      }
    },

    delete: async (id: number): Promise<any> => {
      try {
        const res = await useApiCall("bitacoraForum", ("/post/" + id), {
          method: 'DELETE',
        })
        return res
      } catch (error) {
        console.error("No se han podido obtener todos los posts:", error)
        return null
      }
    },

} as const

type IPostModel = typeof PostModel


export { PostModel, type IPostModel, type PostsFilterType, type PostModelType, type NewPostModel }
