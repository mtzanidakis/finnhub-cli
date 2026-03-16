# finnhub-cli

A command-line client for the [Finnhub](https://finnhub.io/) stock API, written in Go. Outputs JSON to stdout, making it suitable for scripting, pipelines, and AI agent consumption.

## Installation

### From source

```bash
go install github.com/mtzanidakis/finnhub-cli@latest
```

### Build locally

```bash
git clone https://github.com/mtzanidakis/finnhub-cli.git
cd finnhub-cli
make build
```

### Releases

Download pre-built binaries from the [GitHub Releases](https://github.com/mtzanidakis/finnhub-cli/releases) page.

## Configuration

| Variable | Required | Description |
|---|---|---|
| `FINNHUB_API_KEY` | Yes | Your Finnhub API key from [finnhub.io](https://finnhub.io/) |
| `FINNHUB_RATE_LIMIT` | No | `"free"` (default, 30 req/min) or `"premium"` (300 req/min) |

```bash
export FINNHUB_API_KEY="your_api_key_here"
export FINNHUB_RATE_LIMIT="free"
```

## Usage

```
finnhub <command> <subcommand> [flags]
```

Every subcommand accepts `--raw` to produce compact (non-indented) JSON output.

---

### stock -- Stock Market Data

| Subcommand | Description | Example |
|---|---|---|
| `quote` | Real-time quote | `finnhub stock quote --symbol AAPL` |
| `candles` | OHLCV candle data | `finnhub stock candles --symbol AAPL --resolution D --from 2025-01-01 --to 2025-06-01` |
| `profile` | Company profile | `finnhub stock profile --symbol MSFT` |
| `market-status` | Exchange market status | `finnhub stock market-status --exchange US` |
| `symbols` | List symbols on exchange | `finnhub stock symbols --exchange US` |
| `search` | Search for symbols | `finnhub stock search --query "tesla"` |
| `financials` | Reported financials | `finnhub stock financials --symbol AAPL --freq quarterly` |
| `earnings` | Company earnings history | `finnhub stock earnings --symbol AAPL` |

### news -- News & Sentiment

| Subcommand | Description | Example |
|---|---|---|
| `company` | Company-specific news | `finnhub news company --symbol AAPL --from 2025-03-01 --to 2025-03-15` |
| `market` | General market news | `finnhub news market --category general` |
| `sentiment` | News sentiment scores | `finnhub news sentiment --symbol TSLA` |
| `insider-sentiment` | Insider sentiment data | `finnhub news insider-sentiment --symbol MSFT --from 2025-01-01 --to 2025-06-01` |

### technical -- Technical Analysis

| Subcommand | Description | Example |
|---|---|---|
| `indicator` | Technical indicator values | `finnhub technical indicator --symbol AAPL --indicator sma --timeperiod 20` |
| `signals` | Aggregate signal summary | `finnhub technical signals --symbol AAPL --resolution D` |
| `patterns` | Pattern recognition | `finnhub technical patterns --symbol AAPL --resolution D` |
| `support-resistance` | Support/resistance levels | `finnhub technical support-resistance --symbol AAPL` |

### fundamentals -- Fundamental Data

| Subcommand | Description | Example |
|---|---|---|
| `basic` | Basic financial metrics (P/E, margins, etc.) | `finnhub fundamentals basic --symbol AAPL` |
| `reported` | Reported financial statements | `finnhub fundamentals reported --symbol AAPL --freq annual` |
| `sec` | SEC-filed financial statements | `finnhub fundamentals sec --symbol AAPL --freq quarterly` |
| `dividends` | Dividend history | `finnhub fundamentals dividends --symbol AAPL --from 2024-01-01 --to 2025-01-01` |
| `splits` | Stock split history | `finnhub fundamentals splits --symbol AAPL --from 2020-01-01 --to 2025-01-01` |
| `revenue-breakdown` | Revenue segmentation | `finnhub fundamentals revenue-breakdown --symbol AAPL` |

### estimates -- Analyst Estimates

| Subcommand | Description | Example |
|---|---|---|
| `eps` | EPS estimates | `finnhub estimates eps --symbol AAPL` |
| `revenue` | Revenue estimates | `finnhub estimates revenue --symbol MSFT` |
| `ebitda` | EBITDA estimates | `finnhub estimates ebitda --symbol GOOGL` |
| `price-targets` | Price target consensus | `finnhub estimates price-targets --symbol TSLA` |
| `recommendations` | Analyst recommendation trends | `finnhub estimates recommendations --symbol NVDA` |

### ownership -- Ownership Data

| Subcommand | Description | Example |
|---|---|---|
| `insider-transactions` | Insider buy/sell transactions | `finnhub ownership insider-transactions --symbol AAPL` |
| `institutional` | Institutional ownership | `finnhub ownership institutional --symbol AAPL` |
| `portfolio` | Mutual fund holdings by CIK | `finnhub ownership portfolio --cik 0001067983` |
| `congressional` | Congressional trading activity | `finnhub ownership congressional --symbol AAPL --from 2025-01-01 --to 2025-06-01` |

### alternative -- Alternative Data

| Subcommand | Description | Example |
|---|---|---|
| `esg` | ESG (environmental, social, governance) scores | `finnhub alternative esg --symbol AAPL` |
| `social-sentiment` | Social media sentiment | `finnhub alternative social-sentiment --symbol TSLA --from 2025-03-01 --to 2025-03-15` |
| `supply-chain` | Supply chain relationships | `finnhub alternative supply-chain --symbol AAPL` |
| `patents` | US patent filings | `finnhub alternative patents --symbol AAPL --from 2024-01-01 --to 2025-01-01` |

### filings -- SEC Filings

| Subcommand | Description | Example |
|---|---|---|
| `list` | List SEC filings | `finnhub filings list --symbol AAPL` |
| `sentiment` | Filing sentiment analysis | `finnhub filings sentiment --access-number "0000320193-23-000077"` |
| `similarity-index` | Filing similarity index | `finnhub filings similarity-index --symbol AAPL` |

### crypto -- Cryptocurrency

| Subcommand | Description | Example |
|---|---|---|
| `exchanges` | List crypto exchanges | `finnhub crypto exchanges` |
| `symbols` | List symbols on exchange | `finnhub crypto symbols --exchange binance` |
| `profile` | Crypto asset profile | `finnhub crypto profile --symbol BTC` |
| `candles` | Crypto OHLCV candles | `finnhub crypto candles --symbol BINANCE:BTCUSDT --resolution D --from 2025-01-01 --to 2025-03-01` |

### forex -- Foreign Exchange

| Subcommand | Description | Example |
|---|---|---|
| `exchanges` | List forex exchanges | `finnhub forex exchanges` |
| `symbols` | List forex pairs on exchange | `finnhub forex symbols --exchange oanda` |
| `candles` | Forex OHLCV candles | `finnhub forex candles --symbol OANDA:EUR_USD --resolution D` |
| `rates` | Current exchange rates | `finnhub forex rates --base USD` |

### calendar -- Market Calendars

| Subcommand | Description | Example |
|---|---|---|
| `ipo` | Upcoming and recent IPOs | `finnhub calendar ipo --from 2025-03-01 --to 2025-03-31` |
| `earnings` | Earnings calendar | `finnhub calendar earnings --from 2025-03-01 --to 2025-03-07` |
| `economic` | Economic event calendar | `finnhub calendar economic` |
| `fda` | FDA committee meetings | `finnhub calendar fda` |

### events -- Market Events

| Subcommand | Description | Example |
|---|---|---|
| `holidays` | Exchange holiday schedule | `finnhub events holidays --exchange US` |
| `upgrades` | Upgrade/downgrade history | `finnhub events upgrades --symbol AAPL` |
| `mergers` | Merger target countries | `finnhub events mergers` |

---

## Output Format

All commands output JSON to stdout. By default, JSON is pretty-printed with 2-space indentation. Use `--raw` for compact output:

```bash
finnhub stock quote --symbol AAPL --raw
```

Errors are printed to stderr with a non-zero exit code.

## Rate Limiting

The CLI automatically tracks API calls in `~/.finnhub-cli/ratelimit.json` and will pause when the limit is reached. Set `FINNHUB_RATE_LIMIT=premium` if you have a paid plan.

## License

[MIT](LICENSE)
