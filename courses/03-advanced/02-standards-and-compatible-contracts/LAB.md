# Lab: Standards, contracts, and compatible evolution

**Starting point:** your prior capstone commit and today’s prerequisite preparation.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Document success/errors in OpenAPI and notification fields in a JSON schema or equivalent contract.
2. Implement stored idempotency results tied to caller and operation.
3. Reject mismatched payload reuse.
4. Run old-consumer/new-producer compatibility tests for an additive change.

## Acceptance checks

- [ ] Repeated requests create one booking and return a consistent result.
- [ ] Different payload with the same key is rejected.
- [ ] Compatibility expectations are executable or explicitly demonstrated.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
