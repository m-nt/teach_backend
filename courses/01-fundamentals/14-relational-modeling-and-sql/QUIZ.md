# Quiz: Relational modeling and SQL

[Tier index](../README.md) · [Assessment](../../../instructor/assessment.md) · [Capstone](../CAPSTONE.md)

Answer from reasoning; a copied command without an explanation is insufficient. Suggested time: 12 minutes. Total: 6 points.

## Questions

1. **2 points.** In your own words, state the central concept: Relational tables represent entities and relationships. Primary keys identify rows, foreign keys preserve references, and constraints enforce selected business invariants.
2. **2 points.** Scenario: What does a foreign key protect? Explain the expected behavior and why.
3. **2 points.** You changed your artifact and it no longer behaves as intended. Describe the first three debugging checks you would perform.

## Scoring

- 2: correct and explains why.
- 1: partially correct or correct conclusion without a reason.
- 0: incorrect or unsupported.

## Instructor answer key

<details>
<summary>Open after learners submit</summary>

1. A full answer accurately states: Relational tables represent entities and relationships. Primary keys identify rows, foreign keys preserve references, and constraints enforce selected business invariants.
2. It prevents a row from referencing a nonexistent related row.
3. Reproduce the failure, compare the changed input/code with the last known-good version, and read the first relevant error or status; then make one focused change and rerun.

</details>
