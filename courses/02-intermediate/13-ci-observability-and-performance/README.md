# Intermediate 13: CI checks, structured logs, and measured performance

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 180 minutes. **Outcome:** Create a reproducible quality gate and diagnose one request using logs and measurements.

## Prerequisites

Complete [Containerizing Go, configuration, and graceful shutdown](../12-containers-config-and-shutdown/README.md), including its next-class preparation.

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

### 1. A continuous integration gate runs the same checks for every change: formatting, vet, tests, and a build.

Include race checks and database integration tests where the runner supports them. Dependency vulnerability scanning complements tests.

### 2. Structured logs carry fields such as request ID, route, status, and duration.

Do not log passwords, session tokens, OTPs, or unrestricted request bodies. Logs should answer what failed and where without exposing private data.

### 3. A benchmark measures a specific workload on a specific environment.

Record dataset, concurrency, duration, error rate, and latency distribution before changing code. Use profiling to find a bottleneck rather than guessing.

## Exit check

- CI fails for a broken test.
- A request ID links response evidence to logs.
- Measurement includes workload and environment details.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
