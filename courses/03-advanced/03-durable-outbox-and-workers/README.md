# Advanced 03: Transactional outbox and durable notification workers

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 180 minutes. **Outcome:** Commit booking and delivery intent atomically and process work with recoverable claims.

## Prerequisites

Complete [Standards, contracts, and compatible evolution](../02-standards-and-compatible-contracts/README.md), including its next-class preparation.

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

### 1. Committing a booking and then separately enqueueing a message has a failure gap.

An outbox row written in the same database transaction records delivery intent atomically with the booking.

### 2. Workers claim bounded batches with a recoverable lease or transaction-safe locking strategy.

Persist attempt count, next attempt time, and last safe error. A crashed claim must become eligible again.

### 3. At-least-once processing can duplicate delivery.

Give events stable IDs and make internal consumers idempotent using a unique receipt. External email can still duplicate after an ambiguous send result; document the provider’s actual guarantees.

## Exit check

- A committed booking always has delivery intent.
- Abandoned claims become processable again.
- Duplicate events do not duplicate internal business effects.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
