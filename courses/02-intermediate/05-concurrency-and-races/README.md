# Intermediate 05: Goroutines, channels, mutexes, and bounded work

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 180 minutes. **Outcome:** Implement bounded concurrent work and distinguish memory races from business races.

## Prerequisites

Complete [Context, timeouts, defer, and resource lifetimes](../04-context-and-resource-lifetimes/README.md), including its next-class preparation.

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

### 1. A goroutine runs concurrently with other goroutines.

Shared memory needs coordination. A mutex protects an invariant across related reads and writes; locking only the write can still be incorrect.

### 2. Channels communicate values and completion.

The sender normally owns closing a channel. Workers need a stop path; unbounded goroutines and abandoned sends can leak resources.

### 3. The race detector detects certain unsafe memory accesses during executed code.

It does not prove transaction correctness or exercise every schedule. A double booking can be a business race even with no Go memory race.

## Exit check

- No race is reported in exercised paths.
- Cancellation test waits for every worker to stop.
- Worker count remains bounded.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
