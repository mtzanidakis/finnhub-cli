package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/mtzanidakis/finnhub-cli/internal"
)

func runOwnership(args []string) error {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, `Usage: finnhub ownership <subcommand> [flags]

Subcommands:
  insider-transactions    Insider transactions (--symbol)
  institutional           Institutional ownership (--symbol)
  portfolio               Mutual fund holdings (--cik)
  congressional           Congressional trading (--symbol, --from, --to)`)
		return fmt.Errorf("ownership: subcommand required")
	}

	switch args[0] {
	case "insider-transactions":
		return ownershipInsiderTransactions(args[1:])
	case "institutional":
		return ownershipInstitutional(args[1:])
	case "portfolio":
		return ownershipPortfolio(args[1:])
	case "congressional":
		return ownershipCongressional(args[1:])
	default:
		return fmt.Errorf("unknown ownership command %q", args[0])
	}
}

func ownershipInsiderTransactions(args []string) error {
	fs := flag.NewFlagSet("ownership insider-transactions", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	symbol := fs.String("symbol", "", "Stock symbol (required)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *symbol == "" {
		return fmt.Errorf("ownership insider-transactions: --symbol is required")
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.InsiderTransactions(internal.Ctx()).Symbol(*symbol).Execute()
	if err != nil {
		return fmt.Errorf("ownership insider-transactions: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func ownershipInstitutional(args []string) error {
	fs := flag.NewFlagSet("ownership institutional", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	symbol := fs.String("symbol", "", "Stock symbol (required)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *symbol == "" {
		return fmt.Errorf("ownership institutional: --symbol is required")
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.InstitutionalOwnership(internal.Ctx()).Symbol(*symbol).Execute()
	if err != nil {
		return fmt.Errorf("ownership institutional: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func ownershipPortfolio(args []string) error {
	fs := flag.NewFlagSet("ownership portfolio", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	cik := fs.String("cik", "", "CIK number (required)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *cik == "" {
		return fmt.Errorf("ownership portfolio: --cik is required")
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.InstitutionalPortfolio(internal.Ctx()).Cik(*cik).Execute()
	if err != nil {
		return fmt.Errorf("ownership portfolio: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func ownershipCongressional(args []string) error {
	fs := flag.NewFlagSet("ownership congressional", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	symbol := fs.String("symbol", "", "Stock symbol (required)")
	from := fs.String("from", internal.DefaultFrom(), "Start date (YYYY-MM-DD)")
	to := fs.String("to", internal.DefaultTo(), "End date (YYYY-MM-DD)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *symbol == "" {
		return fmt.Errorf("ownership congressional: --symbol is required")
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.CongressionalTrading(internal.Ctx()).Symbol(*symbol).From(*from).To(*to).Execute()
	if err != nil {
		return fmt.Errorf("ownership congressional: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}
