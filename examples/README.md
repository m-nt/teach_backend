# Runnable teaching examples

Each Go directory is an independent module. Run commands inside that directory; there is intentionally no root Go module. These are small teaching slices, not a complete deployed booking product.

| Directory | Lessons | Run / purpose |
|---|---|---|
| [01-hello](01-hello/README.md) | F01–02 | `go run .`: first program |
| [02-booking-cli](02-booking-cli/README.md) | F03–06, F12 | `go run .`, `go test ./...`: entities and rules |
| [03-http-api](03-http-api/README.md) | F07–08 | `go run .`: JSON/HTTP with in-memory reservations |
| [04-sql](04-sql/README.md) | F10 | Schema, synthetic fixtures, and queries |
| [05-db-api](05-db-api/README.md) | F11–13 | PostgreSQL-backed student registration and session catalog |

The first three examples use only the standard library. The database API uses the pgx database/sql adapter. Use a supported installed Go toolchain; `go 1.22` specifies the language baseline, not the preferred current release. The examples listen on loopback by default and have no authentication. Use only synthetic records.

Check all examples with `bash scripts/check-examples.sh` from repository root. See [validation](../instructor/validation.md) for actual executed checks and environmental limitations.
