package internal

import (
	"context"
	"testing"
)

func TestNewClient_WithKey(t *testing.T) {
	t.Setenv("FINNHUB_API_KEY", "test-key-123")

	client, err := NewClient()
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if client == nil {
		t.Fatal("expected non-nil client")
	}
}

func TestNewClient_WithoutKey(t *testing.T) {
	t.Setenv("FINNHUB_API_KEY", "")

	_, err := NewClient()
	if err == nil {
		t.Fatal("expected error when FINNHUB_API_KEY is empty")
	}
	if got := err.Error(); got != "FINNHUB_API_KEY environment variable is required" {
		t.Errorf("unexpected error message: %s", got)
	}
}

func TestCtx(t *testing.T) {
	ctx := Ctx()
	if ctx == nil {
		t.Fatal("expected non-nil context")
	}
	if ctx != context.Background() {
		t.Error("expected context.Background()")
	}
}
