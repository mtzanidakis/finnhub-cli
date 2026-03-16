package internal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// withTempHome sets HOME to a temp dir for the duration of the test,
// so that stateFilePath() resolves to a temporary location.
func withTempHome(t *testing.T) string {
	t.Helper()
	tmp := t.TempDir()
	origHome := os.Getenv("HOME")
	t.Setenv("HOME", tmp)
	t.Cleanup(func() {
		_ = os.Setenv("HOME", origHome)
	})
	return tmp
}

func TestStateFilePath(t *testing.T) {
	tmp := withTempHome(t)
	got := stateFilePath()
	expected := filepath.Join(tmp, ".finnhub-cli", "ratelimit.json")
	if got != expected {
		t.Errorf("expected %s, got %s", expected, got)
	}
}

func TestLoadState_NoFile(t *testing.T) {
	withTempHome(t)
	state, err := loadState()
	if err != nil {
		t.Fatalf("loadState: %v", err)
	}
	if len(state.Timestamps) != 0 {
		t.Errorf("expected empty timestamps, got %v", state.Timestamps)
	}
}

func TestSaveAndLoadState(t *testing.T) {
	withTempHome(t)

	now := time.Now().Unix()
	state := &RateLimitState{
		Timestamps: []int64{now - 10, now - 5, now},
	}

	if err := saveState(state); err != nil {
		t.Fatalf("saveState: %v", err)
	}

	loaded, err := loadState()
	if err != nil {
		t.Fatalf("loadState: %v", err)
	}

	if len(loaded.Timestamps) != len(state.Timestamps) {
		t.Fatalf("expected %d timestamps, got %d", len(state.Timestamps), len(loaded.Timestamps))
	}
	for i, ts := range loaded.Timestamps {
		if ts != state.Timestamps[i] {
			t.Errorf("timestamp[%d]: expected %d, got %d", i, state.Timestamps[i], ts)
		}
	}
}

func TestLoadState_CorruptedFile(t *testing.T) {
	tmp := withTempHome(t)

	dir := filepath.Join(tmp, ".finnhub-cli")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatalf("MkdirAll: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "ratelimit.json"), []byte("not json"), 0o644); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}

	state, err := loadState()
	if err != nil {
		t.Fatalf("loadState should not error on corrupt file, got: %v", err)
	}
	if len(state.Timestamps) != 0 {
		t.Errorf("expected empty timestamps for corrupt file, got %v", state.Timestamps)
	}
}

func TestRateLimit_AddsTimestamp(t *testing.T) {
	withTempHome(t)
	t.Setenv("FINNHUB_RATE_LIMIT", "")

	RateLimit()

	state, err := loadState()
	if err != nil {
		t.Fatalf("loadState: %v", err)
	}
	if len(state.Timestamps) != 1 {
		t.Errorf("expected 1 timestamp after RateLimit(), got %d", len(state.Timestamps))
	}
}

func TestRateLimit_PrunesOldTimestamps(t *testing.T) {
	tmp := withTempHome(t)
	t.Setenv("FINNHUB_RATE_LIMIT", "")

	// Seed state with old timestamps that should be pruned
	old := time.Now().Unix() - 120 // 2 minutes ago
	state := &RateLimitState{
		Timestamps: []int64{old, old + 1, old + 2},
	}

	dir := filepath.Join(tmp, ".finnhub-cli")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatalf("MkdirAll: %v", err)
	}
	data, _ := json.Marshal(state)
	if err := os.WriteFile(filepath.Join(dir, "ratelimit.json"), data, 0o644); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}

	RateLimit()

	loaded, err := loadState()
	if err != nil {
		t.Fatalf("loadState: %v", err)
	}
	// Old timestamps should be pruned; only the new one should remain
	if len(loaded.Timestamps) != 1 {
		t.Errorf("expected 1 timestamp after pruning, got %d: %v", len(loaded.Timestamps), loaded.Timestamps)
	}
}
