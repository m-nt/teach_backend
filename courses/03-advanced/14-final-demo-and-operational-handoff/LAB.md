# Lab: Final demonstration and operational handoff

**Starting point:** your prior capstone commit and today’s prerequisite preparation.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Demonstrate the deployed request and notification flow over verified TLS.
2. Run one randomly selected failure drill and recover.
3. Show restore evidence, quality gates, and ownership/OTP tests.
4. Give the runbook to a peer and answer questions about the guarantees.

## Acceptance checks

- [ ] Mandatory advanced safety and correctness gates pass.
- [ ] A peer can follow the operational runbook.
- [ ] Each student explains a subsystem they personally implemented.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
