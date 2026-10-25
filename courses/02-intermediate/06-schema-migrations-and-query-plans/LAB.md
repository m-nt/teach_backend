# Lab: Schema migrations, indexes, and query plans

**Starting point:** your prior capstone commit and today’s prerequisite preparation.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Write numbered up/down scripts for a nullable column, documenting rollback limits.
2. Apply to a fresh database and to one containing bookings.
3. Generate a larger synthetic dataset using SQL generate_series.
4. Record a query plan before and after an index and explain the observed difference.

## Acceptance checks

- [ ] Existing booking data survives the upgrade.
- [ ] Plan comparison includes row count and query text.
- [ ] Rollback data-loss consequences are stated.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
