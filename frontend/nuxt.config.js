// https://nuxt.com/docs/api/configuration/nuxt-config

import { defineNuxtConfig } from 'nuxt/config'

export default defineNuxtConfig({
  compatibilityDate: "2024-11-01",
  devtools: { enabled: true },
  modules: ["@nuxt/ui", "@nuxtjs/tailwindcss", "@vee-validate/nuxt", "@nuxtjs/apollo"],

  apollo: {
    clients: {
      default: {
        // httpEndpoint: process.env.APOLLO_HTTP_ENDPOINT,
        httpEndpoint:
        process.env.APOLLO_HTTP_ENDPOINT || "http://localhost:8080/v1/graphql",
        // httpEndpoint: 'https://spacex-production.up.railway.app'
        headers: {
          "X-Hasura-Admin-Secret": process.env.HASURA_GRAPHQL_ADMIN_SECRET || "myadminsecretkey",
        },
      },
    },

 
  },


  publicRuntimeConfig: {
    apolloHttpEndpoint: process.env.APOLLO_HTTP_ENDPOINT || "http://localhost:8080/v1/graphql"
  },
});
