export default {
  clientConfigs: {
    default: {
      httpEndpoint:
        process.env.APOLLO_HTTP_ENDPOINT || "http://localhost:8080/v1/graphql",
      // Add other options here, such as authentication headers
    },
  },
};
