export default defineNuxtRouteMiddleware((to) => {

  const error = { statusCode: 404, statusMessage: 'Page Not Found' }

  if (to.path.startsWith('/auth/')) {
    const validParams = ['login', 'register']
    const param = to.params?.auth as string || ''

    if (!validParams.includes(param)) {
      throw createError(error)
    }
  }

})
