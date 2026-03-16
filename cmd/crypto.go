package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/mtzanidakis/finnhub-cli/internal"
)

func runCrypto(args []string) error {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, `Usage: finnhub crypto <subcommand> [flags]

Subcommands:
  exchanges    List crypto exchanges
  symbols      List crypto symbols (--exchange)
  profile      Get crypto profile (--symbol)
  candles      Get crypto candles (--symbol, --resolution, --from, --to)`)
		return fmt.Errorf("crypto: subcommand required")
	}

	switch args[0] {
	case "exchanges":
		return cryptoExchanges(args[1:])
	case "symbols":
		return cryptoSymbols(args[1:])
	case "profile":
		return cryptoProfile(args[1:])
	case "candles":
		return cryptoCandles(args[1:])
	default:
		return fmt.Errorf("unknown crypto command %q", args[0])
	}
}

func cryptoExchanges(args []string) error {
	fs := flag.NewFlagSet("crypto exchanges", flag.ContinueOnError)
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

	result, _, err := client.CryptoExchanges(internal.Ctx()).Execute()
	if err != nil {
		return fmt.Errorf("crypto exchanges: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func cryptoSymbols(args []string) error {
	fs := flag.NewFlagSet("crypto symbols", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	exchange := fs.String("exchange", "", "Exchange code (required)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *exchange == "" {
		return fmt.Errorf("crypto symbols: --exchange is required")
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.CryptoSymbols(internal.Ctx()).Exchange(*exchange).Execute()
	if err != nil {
		return fmt.Errorf("crypto symbols: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func cryptoProfile(args []string) error {
	fs := flag.NewFlagSet("crypto profile", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	symbol := fs.String("symbol", "", "Crypto symbol (required)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *symbol == "" {
		return fmt.Errorf("crypto profile: --symbol is required")
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.CryptoProfile(internal.Ctx()).Symbol(*symbol).Execute()
	if err != nil {
		return fmt.Errorf("crypto profile: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func cryptoCandles(args []string) error {
	fs := flag.NewFlagSet("crypto candles", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	symbol := fs.String("symbol", "", "Crypto symbol (required)")
	resolution := fs.String("resolution", "D", "Candle resolution (1, 5, 15, 30, 60, D, W, M)")
	from := fs.String("from", internal.DefaultFrom(), "Start date (YYYY-MM-DD)")
	to := fs.String("to", internal.DefaultTo(), "End date (YYYY-MM-DD)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *symbol == "" {
		return fmt.Errorf("crypto candles: --symbol is required")
	}

	fromUnix, err := internal.ParseDateUnix(*from)
	if err != nil {
		return fmt.Errorf("crypto candles: --from: %w", err)
	}
	toUnix, err := internal.ParseDateUnix(*to)
	if err != nil {
		return fmt.Errorf("crypto candles: --to: %w", err)
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.CryptoCandles(internal.Ctx()).Symbol(*symbol).Resolution(*resolution).From(fromUnix).To(toUnix).Execute()
	if err != nil {
		return fmt.Errorf("crypto candles: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}
