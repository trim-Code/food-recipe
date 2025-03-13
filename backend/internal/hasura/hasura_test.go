package hasura

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "reflect"
    "testing"
)

func TestExecuteGraphQLQuery(t *testing.T) {
    // Create a test HTTP server
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Check the request method and URL
        if r.Method != http.MethodPost || r.URL.Path != "/v1/graphql" {
            t.Fatalf("Expected POST /v1/graphql, got %s %s", r.Method, r.URL.Path)
        }

        // Check for the required headers
        if r.Header.Get("X-Hasura-Admin-Secret") != "myadminsecretkey" {
            t.Fatalf("Expected X-Hasura-Admin-Secret header to be set")
        }

        // Decode the request body
        var req GraphQLRequest
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            t.Fatalf("Failed to decode request body: %v", err)
        }

        // Check the query and variables
        expectedQuery := "query { test }"
        if req.Query != expectedQuery {
            t.Fatalf("Expected query %q, got %q", expectedQuery, req.Query)
        }

        // Respond with a mock response
        resp := GraphQLResponse{
            Data: json.RawMessage(`{"test": "value"}`),
        }
        if err := json.NewEncoder(w).Encode(resp); err != nil {
            t.Fatalf("Failed to encode response: %v", err)
        }
    }))
    defer server.Close()

    // Override the hasuraEndpoint with the test server URL
    endpoint := server.URL + "/v1/graphql"

    // Set the required headers
    headers := map[string]string{
        "X-Hasura-Admin-Secret": "myadminsecretkey",
    }

    // Execute the GraphQL query
    query := "query { test }"
    variables := map[string]interface{}{}
    resp, err := ExecuteGraphQLQuery(endpoint, headers, query, variables)
    if err != nil {
        t.Fatalf("ExecuteGraphQLQuery failed: %v", err)
    }

    // Check the response data
    var expectedData, actualData map[string]interface{}
    if err := json.Unmarshal([]byte(`{"test": "value"}`), &expectedData); err != nil {
        t.Fatalf("Failed to unmarshal expected data: %v", err)
    }
    if err := json.Unmarshal(resp.Data, &actualData); err != nil {
        t.Fatalf("Failed to unmarshal actual data: %v", err)
    }
    if !reflect.DeepEqual(actualData, expectedData) {
        t.Fatalf("Expected data %v, got %v", expectedData, actualData)
    }
}