# Fundamentals 11: Connecting Go to PostgreSQL

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 150 minutes. **Outcome:** Trace and run a database-backed endpoint with parameterized SQL and explicit error handling.

## Prerequisites

Complete [Relational data modeling and basic SQL](../10-sql-and-data-modeling/README.md), including its next-class preparation.

## Class plan

| Minutes | Activity |
|---|---|
| 00–10 | Retrieval practice and setup check |
| 10–35 | Explain concepts with the whiteboard example |
| 35–55 | Instructor live demonstration |
| 55–65 | Break |
| 65–110 | Guided lab and pair explanation |
| 110–130 | Independent lab change and review |
| 130–140 | Quiz and correction |
| 140–150 | Exit check, homework, next-class preparation |

## Lesson

### 1. A database driver speaks the database protocol; database/sql offers a connection pool API.

Open the pool once for the application and close it at shutdown. Do not create a new pool per request.

### 2. Use placeholders and separate arguments for user input.

Building SQL by concatenating input can change the meaning of the query. Distinguish no matching row from an unavailable database.

### 3. A migration is a versioned schema change.

For today use the supplied schema and seeded database. The example lists and creates students; implementing transaction-safe booking waits for intermediate. This keeps a first database lesson small enough to understand.

## Exit check

- New student survives a Go process restart.
- The query uses arguments, not SQL string concatenation.
- Database failure produces an error response rather than a crash.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
