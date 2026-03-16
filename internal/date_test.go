package internal

import (
	"testing"
	"time"
)

func TestParseDate_Valid(t *testing.T) {
	got, err := ParseDate("2024-03-15")
	if err != nil {
		t.Fatalf("ParseDate: %v", err)
	}
	if got.Year() != 2024 || got.Month() != time.March || got.Day() != 15 {
		t.Errorf("expected 2024-03-15, got %v", got)
	}
}

func TestParseDate_Invalid(t *testing.T) {
	cases := []string{
		"not-a-date",
		"03-15-2024",
		"2024/03/15",
		"",
	}
	for _, tc := range cases {
		_, err := ParseDate(tc)
		if err == nil {
			t.Errorf("ParseDate(%q): expected error, got nil", tc)
		}
	}
}

func TestParseDateUnix_Valid(t *testing.T) {
	ts, err := ParseDateUnix("2024-01-01")
	if err != nil {
		t.Fatalf("ParseDateUnix: %v", err)
	}
	expected, _ := time.Parse("2006-01-02", "2024-01-01")
	if ts != expected.Unix() {
		t.Errorf("expected %d, got %d", expected.Unix(), ts)
	}
}

func TestParseDateUnix_Invalid(t *testing.T) {
	_, err := ParseDateUnix("bad")
	if err == nil {
		t.Fatal("expected error for invalid date")
	}
}

func TestDefaultFrom_Format(t *testing.T) {
	s := DefaultFrom()
	_, err := time.Parse("2006-01-02", s)
	if err != nil {
		t.Fatalf("DefaultFrom returned unparseable date: %q", s)
	}
	expected := time.Now().AddDate(0, 0, -30).Format("2006-01-02")
	if s != expected {
		t.Errorf("expected %s, got %s", expected, s)
	}
}

func TestDefaultTo_Format(t *testing.T) {
	s := DefaultTo()
	_, err := time.Parse("2006-01-02", s)
	if err != nil {
		t.Fatalf("DefaultTo returned unparseable date: %q", s)
	}
	expected := time.Now().Format("2006-01-02")
	if s != expected {
		t.Errorf("expected %s, got %s", expected, s)
	}
}
