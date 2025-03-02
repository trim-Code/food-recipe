package main

import (
	"food-recipe-app/backend/db" // Import the database package
	"food-recipe-app/backend/internal/handlers"
	"food-recipe-app/backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
    // Initialize the database
    db.InitDB()

    r := gin.Default()

    // Public routes (no authentication required)
    r.POST("/signup", handlers.Signup)
    r.POST("/login", handlers.Login)

    // Protected routes (requires authentication)
    r.POST("/create-recipe", middleware.AuthMiddleware(), handlers.CreateRecipe)

    // Start the server
    r.Run(":7001")
}
