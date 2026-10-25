#!/usr/bin/env bash
set -euo pipefail
root=$(cd "$(dirname "$0")/.." && pwd)
export GOCACHE=${GOCACHE:-/tmp/teach-backend-go-build}
export GOMODCACHE=${GOMODCACHE:-/tmp/teach-backend-go-mod}
for example in 01-hello 02-booking-cli 03-http-api 05-db-api; do
  (
    cd "$root/examples/$example"
    printf 'Checking %s\n' "$example"
    unformatted=$(gofmt -l .)
    if [[ -n "$unformatted" ]]; then
      printf 'Run gofmt on:\n%s\n' "$unformatted"
      exit 1
    fi
    go vet ./...
    go test ./...
    go test -race ./...
    go build -o /tmp/teach-backend-"$example" .
  )
done
