# Lab: API contracts, pagination, and middleware

**Starting point:** your prior capstone commit and today’s prerequisite preparation.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Document the two endpoints and their error schemas.
2. Add deterministic pagination with a maximum limit.
3. Add request-ID middleware and JSON error responses.
4. Test malformed bodies, invalid limits, and missing resources.

## Acceptance checks

- [ ] Documented examples match actual responses.
- [ ] Page order is deterministic and size is bounded.
- [ ] Logs include request IDs without request secrets.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
