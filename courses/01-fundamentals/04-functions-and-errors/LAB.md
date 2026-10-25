# Lab: Functions, parameters, and explicit errors

**Starting point:** the referenced repository example or an empty personal exercise folder.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Extract remaining into a function with parameters.
2. Write reserve(capacity, booked int) (int, error), returning an updated booked count.
3. Reject negative inputs as well as full capacity.
4. Call reserve from main, checking err before assigning the returned count.

## Acceptance checks

- [ ] A full class returns an error and keeps the original state.
- [ ] A valid request increments booked once.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
