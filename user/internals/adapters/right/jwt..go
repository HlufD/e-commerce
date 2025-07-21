package adapters

import (
	"fmt"
	"time"

	"github.com/HlufD/users-ms/internals/domain"
	"github.com/golang-jwt/jwt/v5"
)

type JWTAdapter struct {
	secret string
	expiry time.Duration
}

func NewJWTAdapter(secret string, expiry time.Duration) *JWTAdapter {
	return &JWTAdapter{
		secret,
		expiry,
	}
}

func (jw *JWTAdapter) Generate(id string) (string, error) {
	// create a claim
	claims := jwt.MapClaims{
		"sub": id,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(jw.expiry).Unix(),
	}

	// Create a new JWT token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token using the secret key
	signedToken, err := token.SignedString([]byte(jw.secret))

	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signedToken, nil

}

func (jw *JWTAdapter) Validate(tokenString string) (string, error) {
	// Parse the token using the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return jw.secret, nil
	})

	if err != nil {
		return "", fmt.Errorf("token parsing failed: %w", err)
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// Extract and return the user ID (sub) from the claims
		sub, ok := claims["sub"].(string)
		if !ok {
			return "", domain.ErrInvalidToken
		}

		return sub, nil
	}

	return "", domain.ErrInvalidToken
}
