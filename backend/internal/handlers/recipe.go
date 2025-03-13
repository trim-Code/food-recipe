package handlers

import (
	"food-recipe-app/backend/internal/hasura"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func CreateRecipe(c *gin.Context) {
    username, _ := c.Get("username")

    var requestBody struct {
        Title       string `json:"title"`
        Description string `json:"description"`
        Category    string `json:"category"`
        PrepTime    int    `json:"prep_time"`
        Ingredients []struct {
            Name     string `json:"name"`
            Quantity string `json:"quantity"`
        } `json:"ingredients"`
        Steps []struct {
            StepNumber int    `json:"step_number"`
            Instruction string `json:"instruction"`
        } `json:"steps"`
    }

    if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
        return
    }

    query := `
    mutation ($title: String!, $description: String!, $category: String!, $prep_time: Int!, $ingredients: [ingredients_insert_input!]!, $steps: [steps_insert_input!]!) {
        insert_recipes(objects: {title: $title, description: $description, category: $category, prep_time: $prep_time, ingredients: {data: $ingredients}, steps: {data: $steps}}) {
            returning {
                id
            }
        }
    }`
    variables := map[string]interface{}{
        "title":       requestBody.Title,
        "description": requestBody.Description,
        "category":    requestBody.Category,
        "prep_time":   requestBody.PrepTime,
        "ingredients": requestBody.Ingredients,
        "steps":       requestBody.Steps,
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