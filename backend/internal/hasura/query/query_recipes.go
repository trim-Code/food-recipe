package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/machinebox/graphql"
)

// Hasura endpoint & secret
var hasuraURL string
var hasuraAdminSecret string

func init() {
	hasuraURL = os.Getenv("HASURA_GRAPHQL_ENDPOINT")
	if hasuraURL == "" {
		log.Fatal("HASURA_GRAPHQL_ENDPOINT not set")
		return
	}
	hasuraAdminSecret = os.Getenv("HASURA_GRAPHQL_ADMIN_SECRET")
	if hasuraAdminSecret == "" {
		log.Fatal("HASURA_GRAPHQL_ADMIN_SECRET not set")
		return
	}
}

func fetchRecipes() {
	client := graphql.NewClient(hasuraURL)
	req := graphql.NewRequest(`
		query {
			recipes {
				id
				title
				description
				created_at
			}
		}
	`)
	req.Header.Set("X-Hasura-Admin-Secret", hasuraAdminSecret)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var respData struct {
		Recipes []struct {
			ID          string `json:"id"`
			Title       string `json:"title"`
			Description string `json:"description"`
			CreatedAt   string `json:"created_at"`
		}
	}

	if err := client.Run(ctx, req, &respData); err != nil {
		log.Printf("Error fetching recipes: %v", err) // Use Printf instead of Fatal
		return                                          // Return instead of exiting
	}

	if len(respData.Recipes) == 0 {
		fmt.Println("No recipes found")
		return
	}

	for _, recipe := range respData.Recipes {
		fmt.Println("Recipe:", recipe.Title, "-", recipe.Description)
	}
}

func main() {
	fetchRecipes()
}
