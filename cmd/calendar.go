package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/mtzanidakis/finnhub-cli/internal"
)

func runCalendar(args []string) error {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, `Usage: finnhub calendar <subcommand> [flags]

Subcommands:
  ipo          IPO calendar (--from, --to)
  earnings     Earnings calendar (--from, --to)
  economic     Economic calendar
  fda          FDA committee meeting calendar`)
		return fmt.Errorf("calendar: subcommand required")
	}

	switch args[0] {
	case "ipo":
		return calendarIPO(args[1:])
	case "earnings":
		return calendarEarnings(args[1:])
	case "economic":
		return calendarEconomic(args[1:])
	case "fda":
		return calendarFDA(args[1:])
	default:
		return fmt.Errorf("unknown calendar command %q", args[0])
	}
}

func calendarIPO(args []string) error {
	fs := flag.NewFlagSet("calendar ipo", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	from := fs.String("from", internal.DefaultFrom(), "Start date (YYYY-MM-DD)")
	to := fs.String("to", internal.DefaultTo(), "End date (YYYY-MM-DD)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.IpoCalendar(internal.Ctx()).From(*from).To(*to).Execute()
	if err != nil {
		return fmt.Errorf("calendar ipo: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func calendarEarnings(args []string) error {
	fs := flag.NewFlagSet("calendar earnings", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	from := fs.String("from", internal.DefaultFrom(), "Start date (YYYY-MM-DD)")
	to := fs.String("to", internal.DefaultTo(), "End date (YYYY-MM-DD)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.EarningsCalendar(internal.Ctx()).From(*from).To(*to).Execute()
	if err != nil {
		return fmt.Errorf("calendar earnings: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func calendarEconomic(args []string) error {
	fs := flag.NewFlagSet("calendar economic", flag.ContinueOnError)
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

	result, _, err := client.EconomicCalendar(internal.Ctx()).Execute()
	if err != nil {
		return fmt.Errorf("calendar economic: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func calendarFDA(args []string) error {
	fs := flag.NewFlagSet("calendar fda", flag.ContinueOnError)
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

	result, _, err := client.FdaCommitteeMeetingCalendar(internal.Ctx()).Execute()
	if err != nil {
		return fmt.Errorf("calendar fda: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}
