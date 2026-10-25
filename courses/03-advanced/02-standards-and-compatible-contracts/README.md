# Advanced 02: Standards, contracts, and compatible evolution

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 180 minutes. **Outcome:** Version HTTP and event contracts without breaking deployed consumers.

## Prerequisites

Complete [Architecture, requirements, and service boundaries](../01-architecture-and-service-boundaries/README.md), including its next-class preparation.

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

### 1. Use consistent HTTP semantics and a documented error shape.

OpenAPI describes synchronous APIs; a versioned event schema describes asynchronous messages. Standards reduce ambiguity but do not choose business rules for you.

### 2. An event needs an ID, type, schema version, occurrence time, and business reference.

Consumers should tolerate documented additive fields. Removing or changing a field requires a migration strategy.

### 3. Idempotency means a repeated logical operation has the intended single effect.

Define a client key’s scope, request fingerprint, stored response, and expiry. Reusing a key for a different request should be rejected.

## Exit check

- Repeated requests create one booking and return a consistent result.
- Different payload with the same key is rejected.
- Compatibility expectations are executable or explicitly demonstrated.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
