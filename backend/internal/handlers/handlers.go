package handlers

import (
	"food-recipe-app/backend/internal/hasura"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func CreateNewRecipe(c *gin.Context) {
	// Assuming the user is already authenticated, you can access their username from the context
	username, _ := c.Get("username")

	query := `
    mutation ($title: String!, $description: String!) {
        insert_recipes(objects: {title: $title, description: $description}) {
            returning {
                id
            }
        }
    }`
	variables := map[string]interface{}{
		"title":       c.PostForm("title"),
		"description": c.PostForm("description"),
	}

	headers := map[string]string{
		"Content-Type": "application/json",
		"X-Hasura-Admin-Secret": os.Getenv("HASURA_GRAPHQL_ADMIN_SECRET"), // Ensure admin secret is passed
	}
	endpoint := os.Getenv("HASURA_GRAPHQL_ENDPOINT") // Use environment variable for endpoint
	_, err := hasura.ExecuteGraphQLQuery(endpoint, headers, query, variables)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Recipe created successfully",
		// "recipe":  newRecipe,
		"user": username,
	})
}
