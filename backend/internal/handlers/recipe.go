// package handlers

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// // Recipe struct to hold recipe data
// type Recipe struct {
//     Name        string   `json:"name"`
//     Ingredients []string `json:"ingredients"`
//     Instructions string  `json:"instructions"`
// }

// filepath: /Users/mac/Desktop/food-recipe-app/backend/internal/handlers/handlers.go
package handlers

import (
    "net/http"
    "food-recipe-app/backend/internal/hasura"
    "github.com/gin-gonic/gin"
)

func CreateRecipe(c *gin.Context) {
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
    }
    endpoint := "https://your-hasura-endpoint/v1/graphql"
    _, err := hasura.ExecuteGraphQLQuery(endpoint, headers, query, variables)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Recipe created successfully",
        // "recipe":  newRecipe,
        "user":    username,
    })
}