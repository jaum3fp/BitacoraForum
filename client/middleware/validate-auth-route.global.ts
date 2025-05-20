export default defineNuxtRouteMiddleware((to) => {

  if (to.path.startsWith('/auth/')) {
    const validParams = ['login', 'register']
    const param = to.params?.auth as string || ''

    if (!validParams.includes(param)) {
      throw createError({ statusCode: 404, statusMessage: 'Page Not Found' })
    }
  }

})
