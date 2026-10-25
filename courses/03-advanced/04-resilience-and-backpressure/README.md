# Advanced 04: Timeouts, retries, backpressure, and overload

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 180 minutes. **Outcome:** Bound failure amplification and demonstrate recovery from a slow dependency.

## Prerequisites

Complete [Transactional outbox and durable notification workers](../03-durable-outbox-and-workers/README.md), including its next-class preparation.

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

### 1. Timeouts consume an overall request budget.

Retrying every layer multiplies calls; select one retry owner and retry only safe transient operations. Use bounded exponential backoff with jitter.

### 2. Backpressure limits admitted work when capacity is exhausted.

Bound queue depth and worker concurrency, expose lag, and define a clear reject or defer response. More goroutines do not create unlimited downstream capacity.

### 3. A circuit breaker may reduce repeated calls to a failing dependency, but it adds state and recovery behavior.

Compare it against a simpler timeout and concurrency limit before adopting it.

## Exit check

- Permanent failures are not retried indefinitely.
- Cancellation stops retry waits.
- Overload remains bounded and observable.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
