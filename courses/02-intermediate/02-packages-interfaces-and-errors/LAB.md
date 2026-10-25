# Lab: Packages, small interfaces, and error contracts

**Starting point:** your prior capstone commit and today’s prerequisite preparation.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Move startup wiring into cmd/api and the rule into an internal package.
2. Define the smallest storage interface the booking behavior needs.
3. Implement a fake for one test.
4. Map wrapped domain errors to stable HTTP responses without exposing SQL details.

## Acceptance checks

- [ ] Domain package does not import net/http.
- [ ] Fake-driven test proves error mapping behavior.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
