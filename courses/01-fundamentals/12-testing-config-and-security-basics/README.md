# Fundamentals 12: Testing, configuration, and security basics

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 150 minutes. **Outcome:** Write a simple Go test and recognize common unsafe backend behavior.

## Prerequisites

Complete [Connecting Go to PostgreSQL](../11-go-and-postgres/README.md), including its next-class preparation.

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

### 1. A test calls code with controlled inputs and checks the result.

Test a successful booking, a full class, and invalid capacity. Expected results should come from the business rule, not from copying the current output.

### 2. Configuration changes between environments; code should not need editing to change a port or database address.

Read environment variables and keep private values out of Git and logs. An .env file is a convenience, not encryption.

### 3. Authentication asks who the caller is; authorization asks what that caller may do.

The beginner examples have no login and use synthetic data on localhost. Identify password hashing, HTTPS, input limits, parameterized queries, and private logging as concerns without implementing a homemade authentication system.

## Exit check

- The new test fails for the deliberate bug and passes after repair.
- Student distinguishes authentication from authorization.
- No credentials appear in the submitted evidence.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
