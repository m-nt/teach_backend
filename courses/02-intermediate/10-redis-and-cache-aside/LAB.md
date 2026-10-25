# Lab: Redis and cache-aside behavior

**Starting point:** your prior capstone commit and today’s prerequisite preparation.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Define a namespaced catalog key and bounded TTL.
2. Implement cache-aside for a public read.
3. Invalidate after a session update commits.
4. Test hit, miss, expiry, invalidation, and Redis-down fallback.

## Acceptance checks

- [ ] An update becomes visible after invalidation.
- [ ] Booking authority never depends on the cache.
- [ ] Redis outage does not break the public read.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
