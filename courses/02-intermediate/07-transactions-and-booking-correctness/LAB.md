# Lab: Transactions and concurrent booking correctness

**Starting point:** your prior capstone commit and today’s prerequisite preparation.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Implement reservation as one context-aware transaction.
2. Lock the session row before checking current bookings.
3. Rollback on every failure and handle commit failure.
4. Run ten concurrent requests for one remaining seat, then query the database.

## Acceptance checks

- [ ] Exactly one request succeeds and capacity is not exceeded.
- [ ] Duplicate student/session requests are rejected.
- [ ] No external notification runs inside the lock-holding transaction.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
