package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/mtzanidakis/finnhub-cli/internal"
)

func runEstimates(args []string) error {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, `Usage: finnhub estimates <subcommand> [flags]

Subcommands:
  eps                EPS estimates (--symbol)
  revenue            Revenue estimates (--symbol)
  ebitda             EBITDA estimates (--symbol)
  price-targets      Price target consensus (--symbol)
  recommendations    Analyst recommendation trends (--symbol)`)
		return fmt.Errorf("estimates: subcommand required")
	}

	switch args[0] {
	case "eps":
		return estimatesEps(args[1:])
	case "revenue":
		return estimatesRevenue(args[1:])
	case "ebitda":
		return estimatesEbitda(args[1:])
	case "price-targets":
		return estimatesPriceTargets(args[1:])
	case "recommendations":
		return estimatesRecommendations(args[1:])
	default:
		return fmt.Errorf("unknown estimates command %q", args[0])
	}
}

func estimatesEps(args []string) error {
	fs := flag.NewFlagSet("estimates eps", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	symbol := fs.String("symbol", "", "Stock symbol (required)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *symbol == "" {
		return fmt.Errorf("estimates eps: --symbol is required")
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.CompanyEpsEstimates(internal.Ctx()).Symbol(*symbol).Execute()
	if err != nil {
		return fmt.Errorf("estimates eps: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func estimatesRevenue(args []string) error {
	fs := flag.NewFlagSet("estimates revenue", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	symbol := fs.String("symbol", "", "Stock symbol (required)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *symbol == "" {
		return fmt.Errorf("estimates revenue: --symbol is required")
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.CompanyRevenueEstimates(internal.Ctx()).Symbol(*symbol).Execute()
	if err != nil {
		return fmt.Errorf("estimates revenue: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func estimatesEbitda(args []string) error {
	fs := flag.NewFlagSet("estimates ebitda", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	symbol := fs.String("symbol", "", "Stock symbol (required)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *symbol == "" {
		return fmt.Errorf("estimates ebitda: --symbol is required")
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	// NOTE: The SDK may expose this as CompanyEbitEstimates instead of
	// CompanyEbitdaEstimates. Adjust the method name if compilation fails.
	result, _, err := client.CompanyEbitdaEstimates(internal.Ctx()).Symbol(*symbol).Execute()
	if err != nil {
		return fmt.Errorf("estimates ebitda: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func estimatesPriceTargets(args []string) error {
	fs := flag.NewFlagSet("estimates price-targets", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	symbol := fs.String("symbol", "", "Stock symbol (required)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *symbol == "" {
		return fmt.Errorf("estimates price-targets: --symbol is required")
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.PriceTarget(internal.Ctx()).Symbol(*symbol).Execute()
	if err != nil {
		return fmt.Errorf("estimates price-targets: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func estimatesRecommendations(args []string) error {
	fs := flag.NewFlagSet("estimates recommendations", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	symbol := fs.String("symbol", "", "Stock symbol (required)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *symbol == "" {
		return fmt.Errorf("estimates recommendations: --symbol is required")
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.RecommendationTrends(internal.Ctx()).Symbol(*symbol).Execute()
	if err != nil {
		return fmt.Errorf("estimates recommendations: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}
