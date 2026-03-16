package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/mtzanidakis/finnhub-cli/internal"
)

func runTechnical(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: finnhub technical <indicator|signals|patterns|support-resistance> [flags]")
	}

	switch args[0] {
	case "indicator":
		return technicalIndicator(args[1:])
	case "signals":
		return technicalSignals(args[1:])
	case "patterns":
		return technicalPatterns(args[1:])
	case "support-resistance":
		return technicalSupportResistance(args[1:])
	default:
		return fmt.Errorf("unknown command %q", args[0])
	}
}

func technicalIndicator(args []string) error {
	fs := flag.NewFlagSet("finnhub technical indicator", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	symbol := fs.String("symbol", "", "stock symbol (required)")
	resolution := fs.String("resolution", "D", "candle resolution")
	from := fs.String("from", internal.DefaultFrom(), "start date YYYY-MM-DD")
	to := fs.String("to", internal.DefaultTo(), "end date YYYY-MM-DD")
	indicator := fs.String("indicator", "", "indicator name, e.g. sma (required)")
	timeperiod := fs.Int("timeperiod", 14, "indicator time period")
	raw := fs.Bool("raw", false, "compact JSON output")

	if err := fs.Parse(args); err != nil {
		return err
	}

	if *symbol == "" {
		return fmt.Errorf("--symbol is required")
	}
	if *indicator == "" {
		return fmt.Errorf("--indicator is required")
	}

	fromUnix, err := internal.ParseDateUnix(*from)
	if err != nil {
		return err
	}
	toUnix, err := internal.ParseDateUnix(*to)
	if err != nil {
		return err
	}

	internal.RateLimit()

	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	resp, _, err := client.TechnicalIndicator(internal.Ctx()).
		Symbol(*symbol).
		Resolution(*resolution).
		From(fromUnix).
		To(toUnix).
		Indicator(*indicator).
		IndicatorFields(map[string]any{"timeperiod": *timeperiod}).
		Execute()
	if err != nil {
		return fmt.Errorf("API error: %w", err)
	}

	return internal.PrintJSON(resp, *raw)
}

func technicalSignals(args []string) error {
	fs := flag.NewFlagSet("finnhub technical signals", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	symbol := fs.String("symbol", "", "stock symbol (required)")
	resolution := fs.String("resolution", "D", "candle resolution")
	raw := fs.Bool("raw", false, "compact JSON output")

	if err := fs.Parse(args); err != nil {
		return err
	}

	if *symbol == "" {
		return fmt.Errorf("--symbol is required")
	}

	internal.RateLimit()

	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	resp, _, err := client.AggregateIndicator(internal.Ctx()).
		Symbol(*symbol).
		Resolution(*resolution).
		Execute()
	if err != nil {
		return fmt.Errorf("API error: %w", err)
	}

	return internal.PrintJSON(resp, *raw)
}

func technicalPatterns(args []string) error {
	fs := flag.NewFlagSet("finnhub technical patterns", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	symbol := fs.String("symbol", "", "stock symbol (required)")
	resolution := fs.String("resolution", "D", "candle resolution")
	raw := fs.Bool("raw", false, "compact JSON output")

	if err := fs.Parse(args); err != nil {
		return err
	}

	if *symbol == "" {
		return fmt.Errorf("--symbol is required")
	}

	internal.RateLimit()

	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	resp, _, err := client.PatternRecognition(internal.Ctx()).
		Symbol(*symbol).
		Resolution(*resolution).
		Execute()
	if err != nil {
		return fmt.Errorf("API error: %w", err)
	}

	return internal.PrintJSON(resp, *raw)
}

func technicalSupportResistance(args []string) error {
	fs := flag.NewFlagSet("finnhub technical support-resistance", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	symbol := fs.String("symbol", "", "stock symbol (required)")
	resolution := fs.String("resolution", "D", "candle resolution")
	raw := fs.Bool("raw", false, "compact JSON output")

	if err := fs.Parse(args); err != nil {
		return err
	}

	if *symbol == "" {
		return fmt.Errorf("--symbol is required")
	}

	internal.RateLimit()

	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	resp, _, err := client.SupportResistance(internal.Ctx()).
		Symbol(*symbol).
		Resolution(*resolution).
		Execute()
	if err != nil {
		return fmt.Errorf("API error: %w", err)
	}

	return internal.PrintJSON(resp, *raw)
}
