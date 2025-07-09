package main

import (
    "github.com/golang-jwt/jwt/v5"
    "time"
    "errors"
)

var jwtKey = []byte("super_secret_key") // Replace with env var in prod

type Claims struct {
    Username string `json:"username"`
    jwt.RegisteredClaims
}

// GenerateJWT creates a signed token
func GenerateJWT(username string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)

    claims := &Claims{
        Username: username,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

// ValidateJWT checks and returns the user from token
func ValidateJWT(tokenStr string) (string, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil || !token.Valid {
        return "", errors.New("invalid token")
    }

    return claims.Username, nil
}
