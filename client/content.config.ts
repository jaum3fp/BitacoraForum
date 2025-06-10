import { defineCollection, defineContentConfig } from '@nuxt/content'

export default defineContentConfig({
  collections: {
    rules: defineCollection({
      source: 'rules/*',
      type: 'page'
    })
  }
})
