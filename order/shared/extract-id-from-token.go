package shared

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func ExtractIDFromToken(ctx context.Context, token string) (string, error) {
	validateURL := os.Getenv("VALIDATE_TOKEN_URL")
	if validateURL == "" {
		return "", fmt.Errorf("VALIDATE_TOKEN_URL is not set")
	}

	bodyBytes, err := json.Marshal(map[string]string{"token": token})
	if err != nil {
		return "", fmt.Errorf("failed to marshal token: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, validateURL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return "", fmt.Errorf("failed to create token validation request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unauthorized: failed to validate token")
	}
	defer resp.Body.Close()

	var validateResp struct {
		UserID string `json:"userId"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&validateResp); err != nil {
		return "", fmt.Errorf("failed to decode token validation response: %w", err)
	}

	return validateResp.UserID, nil
}
