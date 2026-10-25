# Lab: Timeouts, retries, backpressure, and overload

**Starting point:** your prior capstone commit and today’s prerequisite preparation.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Implement max attempts, jittered delay, and cancellation-aware waits.
2. Add a queue or concurrency limit and observable rejection.
3. Simulate timeout, permanent failure, and recovery.
4. Record total downstream calls and queue lag under the chosen workload.

## Acceptance checks

- [ ] Permanent failures are not retried indefinitely.
- [ ] Cancellation stops retry waits.
- [ ] Overload remains bounded and observable.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
