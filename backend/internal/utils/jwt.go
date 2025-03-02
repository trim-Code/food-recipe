package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key") // Ensure this key is consistent

// Claims struct to define the token payload
type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

// GenerateJWT generates a new JWT token
func GenerateJWT(username string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

// ValidateJWT validates a JWT token and returns the username
// func ValidateJWT(tokenStr string) (string, error) {
//     token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
//         return jwtKey, nil
//     })
//     if err != nil || !token.Valid {
//         return "", err
//     }

//     claims, ok := token.Claims.(*Claims)
//     if !ok {
//         return "", err
//     }

//     return claims.Username, nil
// }


func ValidateJWT(tokenStr string) (string, error) {
    token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil {
        log.Println("JWT Error:", err) // Log error
        return "", err
    }
    if !token.Valid {
        log.Println("Invalid Token")
        return "", fmt.Errorf("invalid token")
    }

    claims, ok := token.Claims.(*Claims)
    if !ok {
        log.Println("Invalid Claims")
        return "", fmt.Errorf("invalid claims")
    }

    return claims.Username, nil
}
