package utils

import (
    "github.com/dgrijalva/jwt-go"
    "time"
)

var jwtSecret = []byte("your_secret_key")

func GenerateToken(username string) (string, error) {
    claims := jwt.MapClaims{
        "username": username,
        "exp":      time.Now().Add(time.Hour * 24).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}
