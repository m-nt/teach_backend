# Intermediate 02: Packages, small interfaces, and error contracts

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 180 minutes. **Outcome:** Separate HTTP, booking behavior, and storage using a small consumer-owned interface.

## Prerequisites

Complete [Go value semantics, pointers, and methods](../01-go-values-and-pointers/README.md), including its next-class preparation.

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

### 1. Package boundaries should make dependencies understandable.

Use cmd/api for startup and internal/booking for behavior only when there is enough code to justify it. Keep the domain independent of HTTP status codes.

### 2. Interfaces describe required behavior and are satisfied implicitly.

Define a small repository interface where it is consumed. A fake makes tests possible, but do not add an interface for every struct.

### 3. Wrap errors with context and preserve identity using %w.

errors.Is identifies a known cause; errors.As extracts a typed error. Decide which errors the HTTP layer maps to 404, 409, or a generic server failure.

## Exit check

- Domain package does not import net/http.
- Fake-driven test proves error mapping behavior.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
