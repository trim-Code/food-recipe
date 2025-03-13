import { defineNuxtPlugin } from "#app";
import {
  ApolloClient,
  InMemoryCache,
  createHttpLink,
} from "@apollo/client/core";
import { DefaultApolloClient } from "@vue/apollo-composable";

export default defineNuxtPlugin((nuxtApp) => {
  const httpLink = createHttpLink({
    uri: process.env.APOLLO_HTTP_ENDPOINT || "http://localhost:8080/v1/graphql",
  });

  const cache = new InMemoryCache();

  const apolloClient = new ApolloClient({
    link: httpLink,
    cache,
  });

  nuxtApp.vueApp.provide(DefaultApolloClient, apolloClient);
});
