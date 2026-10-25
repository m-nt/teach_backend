# Lab: Relational data modeling and basic SQL

**Starting point:** the referenced repository example or an empty personal exercise folder.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Draw tables with primary and foreign keys before running SQL.
2. Apply the supplied schema and seed data.
3. Write SELECT with WHERE and ORDER BY; insert one new student.
4. Join bookings to students and sessions; try a duplicate booking inside a transaction and ROLLBACK after the error.

## Acceptance checks

- [ ] Join returns names and session titles for seeded bookings.
- [ ] Duplicate student/session pair is rejected.
- [ ] Student identifies a primary and a foreign key.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
