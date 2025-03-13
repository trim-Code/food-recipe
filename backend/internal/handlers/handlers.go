// package handlers

// import (
// 	"food-recipe-app/backend/internal/hasura"
// 	"net/http"
// 	"os"

// 	"github.com/gin-gonic/gin"
// )

// func CreateNewRecipe(c *gin.Context) {
// 	// Assuming the user is already authenticated, you can access their username from the context
// 	username, _ := c.Get("username")

// 	var requestBody struct {
// 		Title       string `json:"title"`
// 		Description string `json:"description"`
// 	}

// 	if err := c.ShouldBindJSON(&requestBody); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
// 		return
// 	}

// 	query := `
//     mutation ($title: String!, $description: String!) {
//         insert_recipes(objects: {title: $title, description: $description}) {
//             returning {
//                 id
//             }
//         }
//     }`
// 	variables := map[string]interface{}{
// 		"title":       requestBody.Title,
// 		"description": requestBody.Description,
// 	}

// 	headers := map[string]string{
// 		"Content-Type": "application/json",
// 		"X-Hasura-Admin-Secret": os.Getenv("HASURA_GRAPHQL_ADMIN_SECRET"), // Ensure admin secret is passed
// 	}
// 	endpoint := os.Getenv("HASURA_GRAPHQL_ENDPOINT") // Use environment variable for endpoint
// 	_, err := hasura.ExecuteGraphQLQuery(endpoint, headers, query, variables)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{
// 		"message": "Recipe created successfully",
// 		// "recipe":  newRecipe,
// 		"user": username,
// 	})
// }
package handlers

import (
    "food-recipe-app/backend/internal/hasura"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
)

func CreateNewRecipe(c *gin.Context) {
    username, _ := c.Get("username")

    var requestBody struct {
        Title       string `json:"title"`
        Description string `json:"description"`
    }

    if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
        return
    }

    query := `
    mutation ($title: String!, $description: String!) {
        insert_recipes(objects: {title: $title, description: $description}) {
            returning {
                id
            }
        }
    }`
    variables := map[string]interface{}{
        "title":       requestBody.Title,
        "description": requestBody.Description,
    }

    headers := map[string]string{
        "Content-Type": "application/json",
        "X-Hasura-Admin-Secret": os.Getenv("HASURA_GRAPHQL_ADMIN_SECRET"),
    }
    endpoint := os.Getenv("HASURA_GRAPHQL_ENDPOINT")
    _, err := hasura.ExecuteGraphQLQuery(endpoint, headers, query, variables)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Recipe created successfully",
        "user":    username,
    })
}