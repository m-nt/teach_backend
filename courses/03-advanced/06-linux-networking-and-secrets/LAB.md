# Lab: Linux operations, network boundaries, and secrets

**Starting point:** your prior capstone commit and today’s prerequisite preparation.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Inventory listening ports and the intended callers.
2. Create a least-privilege runtime account on the lab VM.
3. Write and apply a reviewed firewall plan with a tested recovery path.
4. Document secret injection and rotate one synthetic credential.

## Acceptance checks

- [ ] Only intended entrypoints are exposed.
- [ ] Rotation works without a source-code edit.
- [ ] Recovery access is tested before closing the original session.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
