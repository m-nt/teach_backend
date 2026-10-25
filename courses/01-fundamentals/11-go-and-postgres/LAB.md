# Lab: Connecting Go to PostgreSQL

**Starting point:** the referenced repository example or an empty personal exercise folder.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Apply schema and seed data if not already present; start the Go database example.
2. List students and create a new synthetic student via HTTP.
3. Restart the Go process and confirm persistence.
4. Stop db briefly, make a request, inspect the controlled 503 response, then start db again.

## Acceptance checks

- [ ] New student survives a Go process restart.
- [ ] The query uses arguments, not SQL string concatenation.
- [ ] Database failure produces an error response rather than a crash.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
