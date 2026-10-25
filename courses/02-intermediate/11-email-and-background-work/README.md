# Intermediate 11: Email adapters and bounded background work

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 180 minutes. **Outcome:** Send sandbox email through an adapter and explain the limitations of in-memory jobs.

## Prerequisites

Complete [Redis and cache-aside behavior](../10-redis-and-cache-aside/README.md), including its next-class preparation.

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

### 1. Separate notification formatting and delivery from booking rules.

An EmailSender interface lets tests record a message without contacting a provider. Use Mailpit locally and synthetic addresses.

### 2. An in-memory queue can demonstrate asynchronous work but loses queued work on process termination.

Return a response that reflects the real delivery guarantee. Advanced adds a transactional outbox for durable delivery intent.

### 3. Retries need a limit and a distinction between transient and permanent errors.

A timeout after sending may mean the receiver got the message, so blind retries can duplicate delivery.

## Exit check

- Message appears in the local inbox.
- Failed delivery is observable and does not deadlock requests.
- Durability limitations are documented.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
