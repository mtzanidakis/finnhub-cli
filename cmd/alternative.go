package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/mtzanidakis/finnhub-cli/internal"
)

func runAlternative(args []string) error {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, `Usage: finnhub alternative <subcommand> [flags]

Subcommands:
  esg              Get company ESG score (--symbol)
  social-sentiment Get stock social sentiment (--symbol)
  supply-chain     Get supply chain relationships (--symbol)
  patents          Get US patent data (--symbol, --from, --to)`)
		return fmt.Errorf("alternative: subcommand required")
	}

	sub := args[0]
	switch sub {
	case "esg":
		return alternativeEsg(args[1:])
	case "social-sentiment":
		return alternativeSocialSentiment(args[1:])
	case "supply-chain":
		return alternativeSupplyChain(args[1:])
	case "patents":
		return alternativePatents(args[1:])
	default:
		return fmt.Errorf("unknown alternative command %q", sub)
	}
}

func alternativeEsg(args []string) error {
	fs := flag.NewFlagSet("alternative esg", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	symbol := fs.String("symbol", "", "Stock symbol (required)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *symbol == "" {
		return fmt.Errorf("alternative esg: --symbol is required")
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.CompanyEsgScore(internal.Ctx()).Symbol(*symbol).Execute()
	if err != nil {
		return fmt.Errorf("alternative esg: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func alternativeSocialSentiment(args []string) error {
	fs := flag.NewFlagSet("alternative social-sentiment", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	symbol := fs.String("symbol", "", "Stock symbol (required)")
	from := fs.String("from", internal.DefaultFrom(), "Start date (YYYY-MM-DD)")
	to := fs.String("to", internal.DefaultTo(), "End date (YYYY-MM-DD)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *symbol == "" {
		return fmt.Errorf("alternative social-sentiment: --symbol is required")
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.SocialSentiment(internal.Ctx()).Symbol(*symbol).From(*from).To(*to).Execute()
	if err != nil {
		return fmt.Errorf("alternative social-sentiment: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func alternativeSupplyChain(args []string) error {
	fs := flag.NewFlagSet("alternative supply-chain", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	symbol := fs.String("symbol", "", "Stock symbol (required)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *symbol == "" {
		return fmt.Errorf("alternative supply-chain: --symbol is required")
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.SupplyChainRelationships(internal.Ctx()).Symbol(*symbol).Execute()
	if err != nil {
		return fmt.Errorf("alternative supply-chain: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}

func alternativePatents(args []string) error {
	fs := flag.NewFlagSet("alternative patents", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	symbol := fs.String("symbol", "", "Stock symbol (required)")
	from := fs.String("from", internal.DefaultFrom(), "Start date (YYYY-MM-DD)")
	to := fs.String("to", internal.DefaultTo(), "End date (YYYY-MM-DD)")
	raw := fs.Bool("raw", false, "Compact JSON output")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *symbol == "" {
		return fmt.Errorf("alternative patents: --symbol is required")
	}

	internal.RateLimit()
	client, err := internal.NewClient()
	if err != nil {
		return err
	}

	result, _, err := client.StockUsptoPatent(internal.Ctx()).Symbol(*symbol).From(*from).To(*to).Execute()
	if err != nil {
		return fmt.Errorf("alternative patents: %w", err)
	}
	return internal.PrintJSON(result, *raw)
}
