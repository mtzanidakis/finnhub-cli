package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/mtzanidakis/finnhub-cli/internal"
)

func runEvents(args []string) error {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, `Usage: finnhub events <subcommand> [flags]

Subcommands:
  holidays     Market holidays (--exchange)
  upgrades     Upgrade/downgrade history (--symbol)
  mergers      Merger target countries`)
		return fmt.Errorf("events: subcommand required")
	}

	switch args[0] {
	case "holidays":
		return eventsHolidays(args[1:])
	case "upgrades":
		return eventsUpgrades(args[1:])
	case "mergers":
		return eventsMergers(args[1:])
	default:
		return fmt.Errorf("unknown events command %q", args[0])
	}
}

func eventsHolidays(args []string) error {
	fs := flag.NewFlagSet("events holidays", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	exchange := fs.String("exchange", "", "Exchange code (required)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *exchange == "" {
		return fmt.Errorf("events holidays: --exchange is required")
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.MarketHoliday(internal.Ctx()).Exchange(*exchange).Execute()
	if err != nil {
		return fmt.Errorf("events holidays: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func eventsUpgrades(args []string) error {
	fs := flag.NewFlagSet("events upgrades", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	symbol := fs.String("symbol", "", "Stock symbol (required)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *symbol == "" {
		return fmt.Errorf("events upgrades: --symbol is required")
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.UpgradeDowngrade(internal.Ctx()).Symbol(*symbol).Execute()
	if err != nil {
		return fmt.Errorf("events upgrades: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func eventsMergers(args []string) error {
	fs := flag.NewFlagSet("events mergers", flag.ContinueOnError)
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

	result, _, err := client.Country(internal.Ctx()).Execute()
	if err != nil {
		return fmt.Errorf("events mergers: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}
