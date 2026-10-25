# Lab: Architecture, requirements, and service boundaries

**Starting point:** your prior capstone commit and today’s prerequisite preparation.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Write functional requirements and three measurable operational targets.
2. Draw processes, storage ownership, and trust boundaries.
3. Define the notification interface and delivery guarantee.
4. Write an architecture decision explaining the chosen split and its cost.

## Acceptance checks

- [ ] Each table and message has one explicit owner.
- [ ] Failure behavior is stated for every dependency.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
