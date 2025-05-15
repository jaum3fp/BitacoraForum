import { defineCollection, defineContentConfig } from '@nuxt/content'

export default defineContentConfig({
  collections: {
    rules: defineCollection({
      source: 'rules/index.md',
      type: 'page'
    })
  }
})
