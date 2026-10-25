# Intermediate 08: API contracts, pagination, and middleware

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 180 minutes. **Outcome:** Publish a consistent HTTP contract with pagination, validation, and request context.

## Prerequisites

Complete [Transactions and concurrent booking correctness](../07-transactions-and-booking-correctness/README.md), including its next-class preparation.

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

### 1. A contract specifies paths, methods, schemas, statuses, and failure shapes.

OpenAPI can document it, but generated text still needs review. Avoid exposing internal errors as public messages.

### 2. Pagination needs deterministic ordering.

Start with bounded limit/offset for the small course catalog; explain how concurrent changes and large offsets motivate cursor pagination later.

### 3. Middleware handles transport-wide work such as request IDs and logs.

Validate input at the boundary and authorize each protected operation. CORS config controls browser access; it is not an authorization system.

## Exit check

- Documented examples match actual responses.
- Page order is deterministic and size is bounded.
- Logs include request IDs without request secrets.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
