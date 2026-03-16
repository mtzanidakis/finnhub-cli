package internal

import (
	"fmt"
	"time"
)

// ParseDate parses a YYYY-MM-DD date string.
func ParseDate(s string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date %q, expected YYYY-MM-DD", s)
	}
	return t, nil
}

// ParseDateUnix parses a YYYY-MM-DD date string and returns unix timestamp.
func ParseDateUnix(s string) (int64, error) {
	t, err := ParseDate(s)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

// DefaultFrom returns a date string for 30 days ago.
func DefaultFrom() string {
	return time.Now().AddDate(0, 0, -30).Format("2006-01-02")
}

// DefaultTo returns today's date string.
func DefaultTo() string {
	return time.Now().Format("2006-01-02")
}
