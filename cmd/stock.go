package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/mtzanidakis/finnhub-cli/internal"
)

func runStock(args []string) error {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, `Usage: finnhub stock <subcommand> [flags]

Subcommands:
  quote            Get real-time quote (--symbol)
  candles          Get stock candles (--symbol, --resolution, --from, --to)
  profile          Get company profile (--symbol)
  market-status    Get market status (--exchange)
  symbols          List stock symbols (--exchange)
  search           Search symbols (--query)
  financials       Get reported financials (--symbol, --statement, --freq)
  earnings         Get company earnings (--symbol)`)
		return fmt.Errorf("stock: subcommand required")
	}

	sub := args[0]
	switch sub {
	case "quote":
		return stockQuote(args[1:])
	case "candles":
		return stockCandles(args[1:])
	case "profile":
		return stockProfile(args[1:])
	case "market-status":
		return stockMarketStatus(args[1:])
	case "symbols":
		return stockSymbols(args[1:])
	case "search":
		return stockSearch(args[1:])
	case "financials":
		return stockFinancials(args[1:])
	case "earnings":
		return stockEarnings(args[1:])
	default:
		return fmt.Errorf("unknown stock command %q", sub)
	}
}

func stockQuote(args []string) error {
	fs := flag.NewFlagSet("stock quote", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	symbol := fs.String("symbol", "", "Stock symbol (required)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *symbol == "" {
		return fmt.Errorf("stock quote: --symbol is required")
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.Quote(internal.Ctx()).Symbol(*symbol).Execute()
	if err != nil {
		return fmt.Errorf("stock quote: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func stockCandles(args []string) error {
	fs := flag.NewFlagSet("stock candles", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	symbol := fs.String("symbol", "", "Stock symbol (required)")
	resolution := fs.String("resolution", "D", "Candle resolution (1, 5, 15, 30, 60, D, W, M)")
	from := fs.String("from", internal.DefaultFrom(), "Start date (YYYY-MM-DD)")
	to := fs.String("to", internal.DefaultTo(), "End date (YYYY-MM-DD)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *symbol == "" {
		return fmt.Errorf("stock candles: --symbol is required")
	}

	fromUnix, err := internal.ParseDateUnix(*from)
	if err != nil {
		return fmt.Errorf("stock candles: --from: %w", err)
	}
	toUnix, err := internal.ParseDateUnix(*to)
	if err != nil {
		return fmt.Errorf("stock candles: --to: %w", err)
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.StockCandles(internal.Ctx()).Symbol(*symbol).Resolution(*resolution).From(fromUnix).To(toUnix).Execute()
	if err != nil {
		return fmt.Errorf("stock candles: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func stockProfile(args []string) error {
	fs := flag.NewFlagSet("stock profile", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	symbol := fs.String("symbol", "", "Stock symbol (required)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *symbol == "" {
		return fmt.Errorf("stock profile: --symbol is required")
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.CompanyProfile2(internal.Ctx()).Symbol(*symbol).Execute()
	if err != nil {
		return fmt.Errorf("stock profile: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func stockMarketStatus(args []string) error {
	fs := flag.NewFlagSet("stock market-status", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	exchange := fs.String("exchange", "", "Exchange code (required)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *exchange == "" {
		return fmt.Errorf("stock market-status: --exchange is required")
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.MarketStatus(internal.Ctx()).Exchange(*exchange).Execute()
	if err != nil {
		return fmt.Errorf("stock market-status: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func stockSymbols(args []string) error {
	fs := flag.NewFlagSet("stock symbols", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	exchange := fs.String("exchange", "", "Exchange code (required)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *exchange == "" {
		return fmt.Errorf("stock symbols: --exchange is required")
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.StockSymbols(internal.Ctx()).Exchange(*exchange).Execute()
	if err != nil {
		return fmt.Errorf("stock symbols: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func stockSearch(args []string) error {
	fs := flag.NewFlagSet("stock search", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	query := fs.String("query", "", "Search query (required)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *query == "" {
		return fmt.Errorf("stock search: --query is required")
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.SymbolSearch(internal.Ctx()).Q(*query).Execute()
	if err != nil {
		return fmt.Errorf("stock search: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func stockFinancials(args []string) error {
	fs := flag.NewFlagSet("stock financials", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	symbol := fs.String("symbol", "", "Stock symbol (required)")
	freq := fs.String("freq", "annual", "Frequency (annual, quarterly)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *symbol == "" {
		return fmt.Errorf("stock financials: --symbol is required")
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.FinancialsReported(internal.Ctx()).Symbol(*symbol).Freq(*freq).Execute()
	if err != nil {
		return fmt.Errorf("stock financials: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func stockEarnings(args []string) error {
	fs := flag.NewFlagSet("stock earnings", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	symbol := fs.String("symbol", "", "Stock symbol (required)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *symbol == "" {
		return fmt.Errorf("stock earnings: --symbol is required")
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.CompanyEarnings(internal.Ctx()).Symbol(*symbol).Execute()
	if err != nil {
		return fmt.Errorf("stock earnings: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}
