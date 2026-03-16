package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

// PrintJSON marshals v to JSON and prints to stdout.
// If raw is true, output is compact; otherwise pretty-printed.
func PrintJSON(v any, raw bool) error {
	var data []byte
	var err error
	if raw {
		data, err = json.Marshal(v)
	} else {
		data, err = json.MarshalIndent(v, "", "  ")
	}
	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}
	_, err = fmt.Fprintf(os.Stdout, "%s\n", data)
	return err
}

// Fatal prints an error to stderr and exits with code 1.
func Fatal(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "error: "+format+"\n", args...)
	os.Exit(1)
}
