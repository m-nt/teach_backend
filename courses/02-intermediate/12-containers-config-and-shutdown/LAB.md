# Lab: Containerizing Go, configuration, and graceful shutdown

**Starting point:** your prior capstone commit and today’s prerequisite preparation.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Add a multi-stage Dockerfile and .dockerignore.
2. Wire environment variables for db, cache, and mail using service names.
3. Add separate health/readiness endpoints and a shutdown deadline.
4. Rebuild from a clean checkout and perform the SIGTERM exercise.

## Acceptance checks

- [ ] Image starts without local source files.
- [ ] API connects using Compose service names.
- [ ] Shutdown finishes or times out predictably.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
