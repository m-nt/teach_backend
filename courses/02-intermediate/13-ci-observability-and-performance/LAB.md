# Lab: CI checks, structured logs, and measured performance

**Starting point:** your prior capstone commit and today’s prerequisite preparation.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Create a local quality script and wire it into your chosen CI runner.
2. Add structured request logging and safe internal error context.
3. Capture a baseline for one workload and identify the likely bottleneck.
4. Make one small measured improvement or explain why no change is justified.

## Acceptance checks

- [ ] CI fails for a broken test.
- [ ] A request ID links response evidence to logs.
- [ ] Measurement includes workload and environment details.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
