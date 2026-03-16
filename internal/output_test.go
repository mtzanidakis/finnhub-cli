package internal

import (
	"io"
	"os"
	"strings"
	"testing"
)

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe: %v", err)
	}

	origStdout := os.Stdout
	os.Stdout = w

	fn()

	_ = w.Close()
	os.Stdout = origStdout

	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("reading pipe: %v", err)
	}
	return string(out)
}

func TestPrintJSON_Raw(t *testing.T) {
	v := map[string]string{"name": "test", "value": "123"}
	got := captureStdout(t, func() {
		if err := PrintJSON(v, true); err != nil {
			t.Fatalf("PrintJSON: %v", err)
		}
	})

	got = strings.TrimSpace(got)
	// Raw output should be compact (no newlines within the JSON)
	if strings.Contains(got, "\n") {
		t.Errorf("raw output should be compact, got:\n%s", got)
	}
	if !strings.Contains(got, `"name":"test"`) {
		t.Errorf("expected key/value in output, got: %s", got)
	}
}

func TestPrintJSON_Pretty(t *testing.T) {
	v := map[string]string{"key": "val"}
	got := captureStdout(t, func() {
		if err := PrintJSON(v, false); err != nil {
			t.Fatalf("PrintJSON: %v", err)
		}
	})

	// Pretty output should contain indentation
	if !strings.Contains(got, "  ") {
		t.Errorf("pretty output should contain indentation, got:\n%s", got)
	}
	if !strings.Contains(got, `"key": "val"`) {
		t.Errorf("expected key/value in output, got: %s", got)
	}
}

func TestPrintJSON_Unmarshalable(t *testing.T) {
	// Channels cannot be marshaled to JSON
	v := make(chan int)
	err := PrintJSON(v, true)
	if err == nil {
		t.Fatal("expected error for unmarshalable value")
	}
	if !strings.Contains(err.Error(), "json marshal") {
		t.Errorf("expected json marshal error, got: %v", err)
	}
}
