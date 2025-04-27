package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var jwtKey []byte

func init() {
	_ = godotenv.Load()
    secret := os.Getenv("JWT_SECRET_KEY")
    if secret == "" {
        panic("JWT_SECRET_KEY environment variable not set")
    }
    jwtKey = []byte(secret)
}

// GenerateJWT - membuat token JWT dari customer ID
func GenerateJWT(customerID string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "customer_id": customerID,
        "exp":         time.Now().Add(time.Hour * 1).Unix(),
    })

    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

// ValidateJWT - memvalidasi token JWT dan ambil customer_id
func ValidateJWT(tokenString string) (string, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("invalid signing method")
        }
        return jwtKey, nil
    })

    if err != nil {
        return "", err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        customerID, ok := claims["customer_id"].(string)
        if !ok {
            return "", errors.New("invalid token claims")
        }
        return customerID, nil
    }

    return "", errors.New("invalid token")
}
