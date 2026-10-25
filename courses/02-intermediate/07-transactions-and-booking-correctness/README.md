# Intermediate 07: Transactions and concurrent booking correctness

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 180 minutes. **Outcome:** Prevent overbooking under simultaneous requests using a database transaction.

## Prerequisites

Complete [Schema migrations, indexes, and query plans](../06-schema-migrations-and-query-plans/README.md), including its next-class preparation.

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

### 1. A transaction groups changes into one commit or rollback.

Reading availability and inserting later without coordination allows two callers to take the last seat.

### 2. For this project, lock the session row with SELECT FOR UPDATE, count active bookings, enforce capacity, insert, then commit.

Every reservation path must acquire that same lock before deciding. A unique student/session constraint separately prevents duplicate reservations.

### 3. A transaction’s lifetime should be short.

Do not send email while holding the database lock. Return only after commit succeeds, and translate conflicts into the agreed HTTP contract.

## Exit check

- Exactly one request succeeds and capacity is not exceeded.
- Duplicate student/session requests are rejected.
- No external notification runs inside the lock-holding transaction.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
