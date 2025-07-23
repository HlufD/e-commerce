package shared

import (
	"fmt"
	"net/http"
)

func ExtractToken(r *http.Request) (string, error) {

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("authorization token missing")
	}

	const bearerPrefix = "Bearer "
	if len(authHeader) <= len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
		return "", fmt.Errorf("invalid token format")
	}

	token := authHeader[len(bearerPrefix):]
	return token, nil
}
