# finnhub-cli -- AI Agent Usage Guide

## Overview

`finnhub` is a CLI client for the Finnhub stock API. It outputs JSON to stdout and errors to stderr. All commands follow the pattern `finnhub <group> <subcommand> [flags]`.

**Prerequisites:** Set `FINNHUB_API_KEY` environment variable before use.

## Command Reference

All commands accept `--raw` (bool, optional, default: false) for compact JSON output.

### stock

| Subcommand | Flags | Required | Defaults |
|---|---|---|---|
| `quote` | `--symbol` | symbol | -- |
| `candles` | `--symbol`, `--resolution`, `--from`, `--to` | symbol | resolution=D, from=30d ago, to=today |
| `profile` | `--symbol` | symbol | -- |
| `market-status` | `--exchange` | exchange | -- |
| `symbols` | `--exchange` | exchange | -- |
| `search` | `--query` | query | -- |
| `financials` | `--symbol`, `--freq` | symbol | freq=annual |
| `earnings` | `--symbol` | symbol | -- |

### news

| Subcommand | Flags | Required | Defaults |
|---|---|---|---|
| `company` | `--symbol`, `--from`, `--to` | symbol | from=30d ago, to=today |
| `market` | `--category` | -- | category=general (general, forex, crypto, merger) |
| `sentiment` | `--symbol` | symbol | -- |
| `insider-sentiment` | `--symbol`, `--from`, `--to` | symbol | from=30d ago, to=today |

### technical

| Subcommand | Flags | Required | Defaults |
|---|---|---|---|
| `indicator` | `--symbol`, `--indicator`, `--resolution`, `--from`, `--to`, `--timeperiod` | symbol, indicator | resolution=D, from=30d ago, to=today, timeperiod=14 |
| `signals` | `--symbol`, `--resolution` | symbol | resolution=D |
| `patterns` | `--symbol`, `--resolution` | symbol | resolution=D |
| `support-resistance` | `--symbol`, `--resolution` | symbol | resolution=D |

### fundamentals

| Subcommand | Flags | Required | Defaults |
|---|---|---|---|
| `basic` | `--symbol` | symbol | -- |
| `reported` | `--symbol`, `--freq` | symbol | freq=annual |
| `sec` | `--symbol`, `--freq` | symbol | freq=annual |
| `dividends` | `--symbol`, `--from`, `--to` | symbol | from=30d ago, to=today |
| `splits` | `--symbol`, `--from`, `--to` | symbol | from=30d ago, to=today |
| `revenue-breakdown` | `--symbol` | symbol | -- |

### estimates

| Subcommand | Flags | Required | Defaults |
|---|---|---|---|
| `eps` | `--symbol` | symbol | -- |
| `revenue` | `--symbol` | symbol | -- |
| `ebitda` | `--symbol` | symbol | -- |
| `price-targets` | `--symbol` | symbol | -- |
| `recommendations` | `--symbol` | symbol | -- |

### ownership

| Subcommand | Flags | Required | Defaults |
|---|---|---|---|
| `insider-transactions` | `--symbol` | symbol | -- |
| `institutional` | `--symbol` | symbol | -- |
| `portfolio` | `--cik` | cik | -- |
| `congressional` | `--symbol`, `--from`, `--to` | symbol | from=30d ago, to=today |

### alternative

| Subcommand | Flags | Required | Defaults |
|---|---|---|---|
| `esg` | `--symbol` | symbol | -- |
| `social-sentiment` | `--symbol`, `--from`, `--to` | symbol | from=30d ago, to=today |
| `supply-chain` | `--symbol` | symbol | -- |
| `patents` | `--symbol`, `--from`, `--to` | symbol | from=30d ago, to=today |

### filings

| Subcommand | Flags | Required | Defaults |
|---|---|---|---|
| `list` | `--symbol` | symbol | -- |
| `sentiment` | `--access-number` | access-number | -- |
| `similarity-index` | `--symbol` | symbol | -- |

### crypto

| Subcommand | Flags | Required | Defaults |
|---|---|---|---|
| `exchanges` | -- | -- | -- |
| `symbols` | `--exchange` | exchange | -- |
| `profile` | `--symbol` | symbol | -- |
| `candles` | `--symbol`, `--resolution`, `--from`, `--to` | symbol | resolution=D, from=30d ago, to=today |

### forex

| Subcommand | Flags | Required | Defaults |
|---|---|---|---|
| `exchanges` | -- | -- | -- |
| `symbols` | `--exchange` | exchange | -- |
| `candles` | `--symbol`, `--resolution`, `--from`, `--to` | symbol | resolution=D, from=30d ago, to=today |
| `rates` | `--base` | -- | base=USD |

### calendar

| Subcommand | Flags | Required | Defaults |
|---|---|---|---|
| `ipo` | `--from`, `--to` | -- | from=30d ago, to=today |
| `earnings` | `--from`, `--to` | -- | from=30d ago, to=today |
| `economic` | -- | -- | -- |
| `fda` | -- | -- | -- |

### events

| Subcommand | Flags | Required | Defaults |
|---|---|---|---|
| `holidays` | `--exchange` | exchange | -- |
| `upgrades` | `--symbol` | symbol | -- |
| `mergers` | -- | -- | -- |

## Common Workflows

### 1. Analyze a Stock

Get a comprehensive view of a single company:

```bash
finnhub stock quote --symbol AAPL --raw
finnhub stock profile --symbol AAPL --raw
finnhub fundamentals basic --symbol AAPL --raw
finnhub stock earnings --symbol AAPL --raw
finnhub estimates recommendations --symbol AAPL --raw
```

### 2. Compare Earnings Estimates

Compare analyst estimates across multiple symbols:

```bash
finnhub estimates eps --symbol AAPL --raw
finnhub estimates eps --symbol MSFT --raw
finnhub estimates revenue --symbol AAPL --raw
finnhub estimates revenue --symbol MSFT --raw
```

### 3. Market Overview

Get a broad picture of market conditions:

```bash
finnhub news market --category general --raw
finnhub stock market-status --exchange US --raw
finnhub calendar economic --raw
finnhub calendar earnings --from 2025-03-10 --to 2025-03-14 --raw
```

### 4. Crypto Analysis

Explore cryptocurrency markets:

```bash
finnhub crypto exchanges --raw
finnhub crypto symbols --exchange binance --raw
finnhub crypto candles --symbol BINANCE:BTCUSDT --resolution D --from 2025-02-01 --to 2025-03-01 --raw
finnhub crypto profile --symbol BTC --raw
```

## Output Format

- **stdout:** JSON (pretty-printed by default, compact with `--raw`)
- **stderr:** Error messages and rate limit warnings
- **Exit code 0:** Success
- **Exit code 1:** Error (missing flags, API failure, missing API key)

## Rate Limiting

The CLI enforces client-side rate limiting by tracking request timestamps in `~/.finnhub-cli/ratelimit.json`.

| Tier | Limit | Set via |
|---|---|---|
| free (default) | 30 requests/minute | `FINNHUB_RATE_LIMIT=free` or unset |
| premium | 300 requests/minute | `FINNHUB_RATE_LIMIT=premium` |

When the limit is reached, the CLI blocks and prints a warning to stderr until a request slot opens. State persists across invocations.

## Error Handling

Errors follow the format `error: <group> <subcommand>: <details>` on stderr.

Common errors:
- `FINNHUB_API_KEY environment variable is required` -- API key not set
- `--symbol is required` -- missing required flag
- `invalid date "...", expected YYYY-MM-DD` -- malformed date input
- `API error: ...` -- upstream Finnhub API failure

All errors produce exit code 1. Parse stderr for error details when automating.

## Flag Reference

| Flag | Type | Description |
|---|---|---|
| `--symbol` | string | Stock/crypto/forex ticker symbol |
| `--exchange` | string | Exchange code (e.g., US, binance, oanda) |
| `--query` | string | Free-text search string |
| `--resolution` | string | Candle granularity: 1, 5, 15, 30, 60, D, W, M |
| `--from` | string | Start date in YYYY-MM-DD format |
| `--to` | string | End date in YYYY-MM-DD format |
| `--freq` | string | Reporting frequency: annual, quarterly |
| `--indicator` | string | Technical indicator name (e.g., sma, ema, rsi) |
| `--timeperiod` | int | Indicator lookback period (default: 14) |
| `--category` | string | News category: general, forex, crypto, merger |
| `--base` | string | Base currency for forex rates (default: USD) |
| `--cik` | string | SEC CIK number for fund holdings |
| `--access-number` | string | SEC filing access number |
| `--raw` | bool | Compact JSON output (default: false) |
