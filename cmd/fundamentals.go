package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/mtzanidakis/finnhub-cli/internal"
)

func runFundamentals(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: finnhub fundamentals <basic|reported|sec|dividends|splits|revenue-breakdown> [flags]")
	}

	switch args[0] {
	case "basic":
		return fundamentalsBasic(args[1:])
	case "reported":
		return fundamentalsReported(args[1:])
	case "sec":
		return fundamentalsSec(args[1:])
	case "dividends":
		return fundamentalsDividends(args[1:])
	case "splits":
		return fundamentalsSplits(args[1:])
	case "revenue-breakdown":
		return fundamentalsRevenueBreakdown(args[1:])
	default:
		return fmt.Errorf("unknown command %q", args[0])
	}
}

func fundamentalsBasic(args []string) error {
	fs := flag.NewFlagSet("finnhub fundamentals basic", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	symbol := fs.String("symbol", "", "stock symbol (required)")
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

	resp, _, err := client.CompanyBasicFinancials(internal.Ctx()).
		Symbol(*symbol).
		Metric("all").
		Execute()
	if err != nil {
		return fmt.Errorf("API error: %w", err)
	}

	return internal.PrintJSON(resp, *raw)
}

func fundamentalsReported(args []string) error {
	fs := flag.NewFlagSet("finnhub fundamentals reported", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	symbol := fs.String("symbol", "", "stock symbol (required)")
	freq := fs.String("freq", "annual", "reporting frequency (annual or quarterly)")
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

	resp, _, err := client.FinancialsReported(internal.Ctx()).
		Symbol(*symbol).
		Freq(*freq).
		Execute()
	if err != nil {
		return fmt.Errorf("API error: %w", err)
	}

	return internal.PrintJSON(resp, *raw)
}

func fundamentalsSec(args []string) error {
	fs := flag.NewFlagSet("finnhub fundamentals sec", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	symbol := fs.String("symbol", "", "stock symbol (required)")
	freq := fs.String("freq", "annual", "reporting frequency (annual or quarterly)")
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

	resp, _, err := client.FinancialsReported(internal.Ctx()).
		Symbol(*symbol).
		Freq(*freq).
		Execute()
	if err != nil {
		return fmt.Errorf("API error: %w", err)
	}

	return internal.PrintJSON(resp, *raw)
}

func fundamentalsDividends(args []string) error {
	fs := flag.NewFlagSet("finnhub fundamentals dividends", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	symbol := fs.String("symbol", "", "stock symbol (required)")
	from := fs.String("from", internal.DefaultFrom(), "start date YYYY-MM-DD")
	to := fs.String("to", internal.DefaultTo(), "end date YYYY-MM-DD")
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

	resp, _, err := client.StockDividends(internal.Ctx()).
		Symbol(*symbol).
		From(*from).
		To(*to).
		Execute()
	if err != nil {
		return fmt.Errorf("API error: %w", err)
	}

	return internal.PrintJSON(resp, *raw)
}

func fundamentalsSplits(args []string) error {
	fs := flag.NewFlagSet("finnhub fundamentals splits", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	symbol := fs.String("symbol", "", "stock symbol (required)")
	from := fs.String("from", internal.DefaultFrom(), "start date YYYY-MM-DD")
	to := fs.String("to", internal.DefaultTo(), "end date YYYY-MM-DD")
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

	resp, _, err := client.StockSplits(internal.Ctx()).
		Symbol(*symbol).
		From(*from).
		To(*to).
		Execute()
	if err != nil {
		return fmt.Errorf("API error: %w", err)
	}

	return internal.PrintJSON(resp, *raw)
}

func fundamentalsRevenueBreakdown(args []string) error {
	fs := flag.NewFlagSet("finnhub fundamentals revenue-breakdown", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	symbol := fs.String("symbol", "", "stock symbol (required)")
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

	resp, _, err := client.RevenueBreakdown(internal.Ctx()).
		Symbol(*symbol).
		Execute()
	if err != nil {
		return fmt.Errorf("API error: %w", err)
	}

	return internal.PrintJSON(resp, *raw)
}
