# finnhub-cli

Go CLI client for the Finnhub stock API, designed for use by AI agents.

## Build & Test

```bash
go build .              # build binary
go test ./...           # run all tests
go vet ./...            # static analysis
make build              # build via Makefile
make test               # test via Makefile
make lint               # vet + staticcheck
```

## Architecture

- **cmd/**: Cobra command definitions, one file per API group (12 groups, 53 subcommands)
- **internal/**: Shared utilities — client init, output formatting, rate limiting
- Uses `github.com/Finnhub-Stock-API/finnhub-go/v2` official SDK
- Output: JSON to stdout, errors to stderr
- Config via env vars: `FINNHUB_API_KEY` (required), `FINNHUB_OUTPUT_FORMAT`, `FINNHUB_RATE_LIMIT`

## Conventions

- TDD: write tests first, then implementation
- All commands follow the pattern: parse flags → rate limit → API call → JSON output
- Date flags use `YYYY-MM-DD` format, converted to unix timestamps internally for candle endpoints
- Mock the finnhub client interface in tests — no network calls in unit tests
