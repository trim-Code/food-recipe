package hasura

import (
    "bytes"
    "encoding/json"
    "net/http"
    "log"
)

const hasuraEndpoint = "http://localhost:8080/v1/graphql"

type GraphQLRequest struct {
    Query     string                 `json:"query"`
    Variables map[string]interface{} `json:"variables,omitempty"`
}

type GraphQLResponse struct {
    Data   json.RawMessage `json:"data"`
    Errors []struct {
        Message string `json:"message"`
    } `json:"errors"`
}

func ExecuteGraphQLQuery(endpoint string, headers map[string]string, query string, variables map[string]interface{}) (*GraphQLResponse, error) {
    reqBody, err := json.Marshal(GraphQLRequest{Query: query, Variables: variables})
    if err != nil {
        return nil, err
    }

    req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(reqBody))
    if err != nil {
        return nil, err
    }
    req.Header.Set("Content-Type", "application/json")
    for key, value := range headers {
        req.Header.Set(key, value)
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var gqlResp GraphQLResponse
    if err := json.NewDecoder(resp.Body).Decode(&gqlResp); err != nil {
        return nil, err
    }

    if len(gqlResp.Errors) > 0 {
        log.Println("GraphQL errors:", gqlResp.Errors)
    }

    return &gqlResp, nil
}