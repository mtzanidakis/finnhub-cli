package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/mtzanidakis/finnhub-cli/internal"
)

func runForex(args []string) error {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, `Usage: finnhub forex <subcommand> [flags]

Subcommands:
  exchanges    List forex exchanges
  symbols      List forex symbols (--exchange)
  candles      Get forex candles (--symbol, --resolution, --from, --to)
  rates        Get forex rates (--base)`)
		return fmt.Errorf("forex: subcommand required")
	}

	switch args[0] {
	case "exchanges":
		return forexExchanges(args[1:])
	case "symbols":
		return forexSymbols(args[1:])
	case "candles":
		return forexCandles(args[1:])
	case "rates":
		return forexRates(args[1:])
	default:
		return fmt.Errorf("unknown forex command %q", args[0])
	}
}

func forexExchanges(args []string) error {
	fs := flag.NewFlagSet("forex exchanges", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.ForexExchanges(internal.Ctx()).Execute()
	if err != nil {
		return fmt.Errorf("forex exchanges: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func forexSymbols(args []string) error {
	fs := flag.NewFlagSet("forex symbols", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	exchange := fs.String("exchange", "", "Exchange code (required)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *exchange == "" {
		return fmt.Errorf("forex symbols: --exchange is required")
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.ForexSymbols(internal.Ctx()).Exchange(*exchange).Execute()
	if err != nil {
		return fmt.Errorf("forex symbols: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func forexCandles(args []string) error {
	fs := flag.NewFlagSet("forex candles", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	symbol := fs.String("symbol", "", "Forex symbol (required)")
	resolution := fs.String("resolution", "D", "Candle resolution (1, 5, 15, 30, 60, D, W, M)")
	from := fs.String("from", internal.DefaultFrom(), "Start date (YYYY-MM-DD)")
	to := fs.String("to", internal.DefaultTo(), "End date (YYYY-MM-DD)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *symbol == "" {
		return fmt.Errorf("forex candles: --symbol is required")
	}

	fromUnix, err := internal.ParseDateUnix(*from)
	if err != nil {
		return fmt.Errorf("forex candles: --from: %w", err)
	}
	toUnix, err := internal.ParseDateUnix(*to)
	if err != nil {
		return fmt.Errorf("forex candles: --to: %w", err)
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.ForexCandles(internal.Ctx()).Symbol(*symbol).Resolution(*resolution).From(fromUnix).To(toUnix).Execute()
	if err != nil {
		return fmt.Errorf("forex candles: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func forexRates(args []string) error {
	fs := flag.NewFlagSet("forex rates", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	base := fs.String("base", "USD", "Base currency")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.ForexRates(internal.Ctx()).Base(*base).Execute()
	if err != nil {
		return fmt.Errorf("forex rates: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}
