# Lab: Delivery pipelines, release rollback, and safe migrations

**Starting point:** your prior capstone commit and today’s prerequisite preparation.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Build and record a versioned image digest.
2. Create a pipeline with tests, image build, and an explicit promotion step.
3. Deploy an additive change and run a synthetic booking check.
4. Roll back the app and verify old code still works with the schema.

## Acceptance checks

- [ ] Release evidence names a digest and source revision.
- [ ] Rollback completes without losing bookings.
- [ ] Migration ownership is explicit.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
