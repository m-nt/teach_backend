# Lab: Go value semantics, pointers, and methods

**Starting point:** your prior capstone commit and today’s prerequisite preparation.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Predict outputs for a copied Session and a pointer to it.
2. Add a Reserve method with a pointer receiver and error return.
3. Copy a slice using make/copy and prove the copy is independent.
4. Explain who owns the mutable session value in your capstone.

## Acceptance checks

- [ ] Prediction matches observed struct and slice behavior.
- [ ] Tests prove Reserve mutates only the intended value.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
