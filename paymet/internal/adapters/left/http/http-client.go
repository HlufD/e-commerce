package adapters

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type HttpClient struct {
	client  *http.Client
	baseURL string
}

func NewHttpClient(baseURL string, timeout time.Duration) *HttpClient {
	return &HttpClient{
		client: &http.Client{
			Timeout: timeout,
		},
		baseURL: baseURL,
	}
}

func (hc *HttpClient) Get(ctx context.Context, path string, result any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, hc.baseURL+path, nil)
	if err != nil {
		return err
	}

	resp, err := hc.client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return parseResponse(resp, result)
}

func (hc *HttpClient) Post(ctx context.Context, path string, body any, result any) error {
	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, hc.baseURL+path, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := hc.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return parseResponse(resp, result)
}

func parseResponse(resp *http.Response, result any) error {
	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return errors.New(string(body))
	}

	if result == nil {
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(result)
}

func (hc *HttpClient) Put(ctx context.Context, path string, body any, result any) error {
	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, hc.baseURL+path, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := hc.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return parseResponse(resp, result)
}
