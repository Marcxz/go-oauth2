package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("JesusManuelCuen")

type CustomClaims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Company  string `json:"company"`
	jwt.RegisteredClaims
}

// GenerateJWT - generates a new token JWT
func GenerateJWT(username, email, company string) (string, error) {
	claims := CustomClaims{
		Username: username,
		Email:    email,
		Company:  company,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
			Issuer:    "go-oauth2",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWT - validates the token
func ValidateJWT(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Validar el token y extraer los claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
