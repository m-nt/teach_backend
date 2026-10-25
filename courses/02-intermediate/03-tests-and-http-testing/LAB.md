# Lab: Table-driven tests, HTTP tests, and integration boundaries

**Starting point:** your prior capstone commit and today’s prerequisite preparation.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Write four table cases for Reserve.
2. Write an HTTP test for malformed JSON and another for a successful request.
3. Add a real PostgreSQL test using a disposable transaction or test database.
4. Run tests twice and in a different order to reveal shared-state assumptions.

## Acceptance checks

- [ ] Tests fail for a deliberate wrong status or extra booking.
- [ ] Integration fixture does not modify classmates’ data.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
