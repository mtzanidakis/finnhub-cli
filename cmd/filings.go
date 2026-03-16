package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/mtzanidakis/finnhub-cli/internal"
)

func runFilings(args []string) error {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, `Usage: finnhub filings <subcommand> [flags]

Subcommands:
  list              List SEC filings (--symbol)
  sentiment         Get filing sentiment (--access-number)
  similarity-index  Get similarity index (--symbol)`)
		return fmt.Errorf("filings: subcommand required")
	}

	sub := args[0]
	switch sub {
	case "list":
		return filingsList(args[1:])
	case "sentiment":
		return filingsSentiment(args[1:])
	case "similarity-index":
		return filingsSimilarityIndex(args[1:])
	default:
		return fmt.Errorf("unknown filings command %q", sub)
	}
}

func filingsList(args []string) error {
	fs := flag.NewFlagSet("filings list", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	symbol := fs.String("symbol", "", "Stock symbol (required)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *symbol == "" {
		return fmt.Errorf("filings list: --symbol is required")
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.Filings(internal.Ctx()).Symbol(*symbol).Execute()
	if err != nil {
		return fmt.Errorf("filings list: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func filingsSentiment(args []string) error {
	fs := flag.NewFlagSet("filings sentiment", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	accessNumber := fs.String("access-number", "", "Filing access number (required)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *accessNumber == "" {
		return fmt.Errorf("filings sentiment: --access-number is required")
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.FilingsSentiment(internal.Ctx()).AccessNumber(*accessNumber).Execute()
	if err != nil {
		return fmt.Errorf("filings sentiment: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func filingsSimilarityIndex(args []string) error {
	fs := flag.NewFlagSet("filings similarity-index", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	symbol := fs.String("symbol", "", "Stock symbol (required)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *symbol == "" {
		return fmt.Errorf("filings similarity-index: --symbol is required")
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.SimilarityIndex(internal.Ctx()).Symbol(*symbol).Execute()
	if err != nil {
		return fmt.Errorf("filings similarity-index: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}
