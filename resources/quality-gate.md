# Repeatable quality gate

For supplied examples, run from repository root:

```sh
bash scripts/check-examples.sh
python3 scripts/check-curriculum.py
```

For a student application, start with this script in its module root and adapt the integration setup to the CI runner:

```sh
#!/usr/bin/env sh
set -eu
unformatted=$(gofmt -l .)
if [ -n "$unformatted" ]; then
  printf '%s
' "$unformatted"
  exit 1
fi
go vet ./...
go test ./...
go test -race ./...
go build ./...
```

The race detector requires a supported platform/toolchain and normally a working C compiler. Record a skipped check explicitly. Add database integration tests against an isolated test database, then a dependency vulnerability check using the official Go security tooling. Pin the tool version in the real pipeline; do not claim that compilation is a vulnerability audit.

Provider-neutral CI stages: checkout → install selected supported Go → restore dependency cache → download verified dependencies → format/vet/unit/race → start isolated database and apply migrations → integration checks → build versioned image → record digest → publish artifact to the chosen registry. Deployment promotion should follow the team’s chosen release process. The repository does not assume GitHub, GitLab, or a paid runner.

## Measurement worksheet

Record source revision, toolchain, CPU/RAM, dataset size, request mix, concurrency, run duration, error rate, throughput, and latency percentiles. Compare the same workload before/after a change. Explain warmup and cache state. Do not claim universal speedups from one laptop run.
