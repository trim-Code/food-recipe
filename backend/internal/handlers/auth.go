package handlers

import (
	"fmt"
	"food-recipe-app/backend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// User struct to store username and hashed password
type User struct {
	Username string
	Password string
}

// Simulated user storage (Replace this with a DB in production)
var users = map[string]User{}

// Login handler to authenticate users
// Login handler to authenticate users
func Login(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Bind JSON input
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Check if user exists
	user, exists := users[credentials.Username]
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Debugging: Print stored and entered password
	fmt.Println("Stored Hashed Password:", user.Password)
	fmt.Println("Entered Password:", credentials.Password)

	// Compare password hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		fmt.Println("Password Mismatch:", err) // Debugging
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(credentials.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Return the token
	c.JSON(http.StatusOK, gin.H{"token": token})
}


// Signup handler to register new users
func Signup(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Bind JSON input
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Check if user already exists
	if _, exists := users[credentials.Username]; exists {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}

	// Store the user (Replace this with a DB insert in production)
	users[credentials.Username] = User{
		Username: credentials.Username,
		Password: string(hashedPassword),
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(credentials.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "token": token})
}
