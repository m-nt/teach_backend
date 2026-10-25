# Lab: Context, timeouts, defer, and resource lifetimes

**Starting point:** your prior capstone commit and today’s prerequisite preparation.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Pass request context through service to repository.
2. Set a bounded timeout for a dependency call.
3. Close bodies and rows on all paths and check iteration errors.
4. Write a cancellation test using a controlled signal rather than a long sleep.

## Acceptance checks

- [ ] Cancellation reaches the fake dependency.
- [ ] Timeout produces a controlled response and no extra mutation.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
