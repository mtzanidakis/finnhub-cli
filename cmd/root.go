package cmd

import (
	"fmt"
	"os"
)

var groups = map[string]func([]string) error{
	"stock":        runStock,
	"news":         runNews,
	"technical":    runTechnical,
	"fundamentals": runFundamentals,
	"estimates":    runEstimates,
	"ownership":    runOwnership,
	"alternative":  runAlternative,
	"filings":      runFilings,
	"crypto":       runCrypto,
	"forex":        runForex,
	"calendar":     runCalendar,
	"events":       runEvents,
}

// Run dispatches to the appropriate command group.
func Run(args []string) error {
	group := args[0]

	if group == "help" || group == "--help" || group == "-h" {
		Usage()
		return nil
	}

	fn, ok := groups[group]
	if !ok {
		return fmt.Errorf("unknown command %q, run 'finnhub help' for usage", group)
	}

	subArgs := args[1:]
	return fn(subArgs)
}

// Usage prints help text to stderr.
func Usage() {
	fmt.Fprintln(os.Stderr, `Usage: finnhub <command> <subcommand> [flags]

Commands:
  stock          Stock market data (quote, candles, profile, symbols, search, market-status, financials, earnings)
  news           News & sentiment (company, market, sentiment, insider-sentiment)
  technical      Technical analysis (indicator, signals, patterns, support-resistance)
  fundamentals   Fundamentals (basic, reported, sec, dividends, splits, revenue-breakdown)
  estimates      Estimates (eps, revenue, ebitda, price-targets, recommendations)
  ownership      Ownership (insider-transactions, institutional, portfolio, congressional)
  alternative    Alternative data (esg, social-sentiment, supply-chain, patents)
  filings        SEC filings (list, sentiment, similarity-index)
  crypto         Crypto (exchanges, symbols, profile, candles)
  forex          Forex (exchanges, symbols, candles, rates)
  calendar       Calendar (ipo, earnings, economic, fda)
  events         Market events (holidays, upgrades, mergers)

Environment:
  FINNHUB_API_KEY       Required. Your Finnhub API key.
  FINNHUB_RATE_LIMIT    Optional. "free" (default, 30 rpm) or "premium" (300 rpm).

Flags:
  --raw                 Compact JSON output (no indentation).`)
}
