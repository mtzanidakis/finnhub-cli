package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// RateLimitState persists request timestamps across CLI invocations.
type RateLimitState struct {
	Timestamps []int64 `json:"timestamps"`
}

func stateFilePath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".finnhub-cli", "ratelimit.json")
}

func loadState() (*RateLimitState, error) {
	path := stateFilePath()
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &RateLimitState{}, nil
		}
		return nil, err
	}
	var state RateLimitState
	if err := json.Unmarshal(data, &state); err != nil {
		return &RateLimitState{}, nil
	}
	return &state, nil
}

func saveState(state *RateLimitState) error {
	path := stateFilePath()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	data, err := json.Marshal(state)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

// RateLimit checks the persistent rate limit state and waits if necessary.
func RateLimit() {
	tier := os.Getenv("FINNHUB_RATE_LIMIT")
	limit := 30 // free tier
	if tier == "premium" {
		limit = 300
	}

	state, err := loadState()
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: could not load rate limit state: %v\n", err)
		return
	}

	now := time.Now().Unix()
	windowStart := now - 60

	// prune old timestamps
	var recent []int64
	for _, ts := range state.Timestamps {
		if ts > windowStart {
			recent = append(recent, ts)
		}
	}

	if len(recent) >= limit {
		// wait until the oldest timestamp in the window expires
		waitUntil := recent[0] + 60
		waitDur := time.Duration(waitUntil-now) * time.Second
		if waitDur > 0 {
			fmt.Fprintf(os.Stderr, "rate limit reached, waiting %v...\n", waitDur)
			time.Sleep(waitDur)
		}
		// prune again after waiting
		now = time.Now().Unix()
		windowStart = now - 60
		recent = nil
		for _, ts := range state.Timestamps {
			if ts > windowStart {
				recent = append(recent, ts)
			}
		}
	}

	recent = append(recent, time.Now().Unix())
	state.Timestamps = recent
	if err := saveState(state); err != nil {
		fmt.Fprintf(os.Stderr, "warning: could not save rate limit state: %v\n", err)
	}
}
