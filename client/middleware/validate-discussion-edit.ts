import { PostModel } from "~/models/post"

export default defineNuxtRouteMiddleware(async (to) => {
  const userStore = useUserStore()
  const postId = parseInt(to.params.id as string)
  const post = await PostModel.getPostById(postId)

  if (post.owner_id !== userStore.user?.id) {
    throw createError({
      statusCode: 403,
      statusMessage: 'No tienes permiso para editar este post',
      fatal: true
    })
  }
})
