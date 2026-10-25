# Advanced 01: Architecture, requirements, and service boundaries

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 180 minutes. **Outcome:** Define a small multi-service system with explicit ownership and failure assumptions.

## Prerequisites

Pass this tier’s entry gate and bring the prior capstone.

## Class plan

| Minutes | Activity |
|---|---|
| 00–15 | Retrieval practice and prerequisite check |
| 15–45 | Concepts, design discussion, and whiteboard |
| 45–70 | Live demonstration and failure case |
| 70–80 | Break |
| 80–135 | Guided implementation lab |
| 135–160 | Failure tests and peer review |
| 160–170 | Quiz and correction |
| 170–180 | Exit check and next-class preparation |

## Lesson

### 1. Start from requirements: reserve a seat, notify the student, and operate the service when dependencies fail.

A service boundary should represent ownership and a reason to deploy independently, not merely a table.

### 2. Keep booking and identity in one business API initially; extract a notification worker/service that owns delivery attempts.

PostgreSQL owns durable business state, Redis owns disposable cache and bounded ephemeral state, and the mail adapter contacts an SMTP sandbox or provider.

### 3. A network boundary adds latency and partial failure.

Record who owns each table, what callers may assume, and what happens when the notification service is down. A modular monolith remains a valid choice for domains that do not need a split.

## Exit check

- Each table and message has one explicit owner.
- Failure behavior is stated for every dependency.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
