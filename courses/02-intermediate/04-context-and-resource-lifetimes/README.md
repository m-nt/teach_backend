# Intermediate 04: Context, timeouts, defer, and resource lifetimes

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 180 minutes. **Outcome:** Propagate cancellation and close resources without leaking work.

## Prerequisites

Complete [Table-driven tests, HTTP tests, and integration boundaries](../03-tests-and-http-testing/README.md), including its next-class preparation.

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

### 1. A context carries cancellation, deadlines, and request-scoped values across calls.

Pass it explicitly as the first parameter where appropriate. A request’s context should reach database and outbound HTTP operations.

### 2. A timeout limits waiting only when the called operation observes cancellation.

Always invoke the cancel function returned by WithTimeout, usually using defer. Do not replace request context with Background deep inside the call chain.

### 3. defer runs when the surrounding function returns.

Close response bodies and rows, check rows.Err, and avoid accumulating deferred cleanup inside a long-lived loop.

## Exit check

- Cancellation reaches the fake dependency.
- Timeout produces a controlled response and no extra mutation.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
