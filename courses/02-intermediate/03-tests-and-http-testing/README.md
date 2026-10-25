# Intermediate 03: Table-driven tests, HTTP tests, and integration boundaries

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 180 minutes. **Outcome:** Test business behavior and handlers with controlled fixtures and useful assertions.

## Prerequisites

Complete [Packages, small interfaces, and error contracts](../02-packages-interfaces-and-errors/README.md), including its next-class preparation.

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

### 1. Table-driven tests describe cases using inputs and expected outcomes.

Name cases after behavior, especially boundaries and failures. An assertion should fail if the requirement is broken.

### 2. httptest records responses without a public server.

Test status, response shape, and the absence of unintended state changes. Avoid asserting incidental whitespace or private implementation details.

### 3. Unit tests isolate rules; integration tests verify the real database boundary.

Keep database test fixtures isolated and cleanup explicit so test order does not matter.

## Exit check

- Tests fail for a deliberate wrong status or extra booking.
- Integration fixture does not modify classmates’ data.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
