package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIClient struct {
	BaseURL string
}

type APIResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func (client *APIClient) GetData() (*APIResponse, error) {
	resp, err := http.Get(client.BaseURL + "/data")
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &apiResp, nil
}
