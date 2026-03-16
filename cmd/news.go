package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/mtzanidakis/finnhub-cli/internal"
)

func runNews(args []string) error {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, `Usage: finnhub news <subcommand> [flags]

Subcommands:
  company            Get company news (--symbol, --from, --to)
  market             Get market news (--category)
  sentiment          Get news sentiment (--symbol)
  insider-sentiment  Get insider sentiment (--symbol, --from, --to)`)
		return fmt.Errorf("news: subcommand required")
	}

	sub := args[0]
	switch sub {
	case "company":
		return newsCompany(args[1:])
	case "market":
		return newsMarket(args[1:])
	case "sentiment":
		return newsSentiment(args[1:])
	case "insider-sentiment":
		return newsInsiderSentiment(args[1:])
	default:
		return fmt.Errorf("unknown news command %q", sub)
	}
}

func newsCompany(args []string) error {
	fs := flag.NewFlagSet("news company", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	symbol := fs.String("symbol", "", "Stock symbol (required)")
	from := fs.String("from", internal.DefaultFrom(), "Start date (YYYY-MM-DD)")
	to := fs.String("to", internal.DefaultTo(), "End date (YYYY-MM-DD)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *symbol == "" {
		return fmt.Errorf("news company: --symbol is required")
	}

	if _, err := internal.ParseDate(*from); err != nil {
		return fmt.Errorf("news company: --from: %w", err)
	}
	if _, err := internal.ParseDate(*to); err != nil {
		return fmt.Errorf("news company: --to: %w", err)
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.CompanyNews(internal.Ctx()).Symbol(*symbol).From(*from).To(*to).Execute()
	if err != nil {
		return fmt.Errorf("news company: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func newsMarket(args []string) error {
	fs := flag.NewFlagSet("news market", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	category := fs.String("category", "general", "News category (general, forex, crypto, merger)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.MarketNews(internal.Ctx()).Category(*category).Execute()
	if err != nil {
		return fmt.Errorf("news market: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func newsSentiment(args []string) error {
	fs := flag.NewFlagSet("news sentiment", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	symbol := fs.String("symbol", "", "Stock symbol (required)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *symbol == "" {
		return fmt.Errorf("news sentiment: --symbol is required")
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.NewsSentiment(internal.Ctx()).Symbol(*symbol).Execute()
	if err != nil {
		return fmt.Errorf("news sentiment: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func newsInsiderSentiment(args []string) error {
	fs := flag.NewFlagSet("news insider-sentiment", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	symbol := fs.String("symbol", "", "Stock symbol (required)")
	from := fs.String("from", internal.DefaultFrom(), "Start date (YYYY-MM-DD)")
	to := fs.String("to", internal.DefaultTo(), "End date (YYYY-MM-DD)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *symbol == "" {
		return fmt.Errorf("news insider-sentiment: --symbol is required")
	}

	if _, err := internal.ParseDate(*from); err != nil {
		return fmt.Errorf("news insider-sentiment: --from: %w", err)
	}
	if _, err := internal.ParseDate(*to); err != nil {
		return fmt.Errorf("news insider-sentiment: --to: %w", err)
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.InsiderSentiment(internal.Ctx()).Symbol(*symbol).From(*from).To(*to).Execute()
	if err != nil {
		return fmt.Errorf("news insider-sentiment: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}
