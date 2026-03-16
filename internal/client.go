package internal

import (
	"context"
	"fmt"
	"os"

	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
)

// NewClient creates a finnhub API client configured from FINNHUB_API_KEY.
func NewClient() (*finnhub.DefaultApiService, error) {
	key := os.Getenv("FINNHUB_API_KEY")
	if key == "" {
		return nil, fmt.Errorf("FINNHUB_API_KEY environment variable is required")
	}

	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", key)
	client := finnhub.NewAPIClient(cfg)
	return client.DefaultApi, nil
}

// Ctx returns a background context.
func Ctx() context.Context {
	return context.Background()
}
