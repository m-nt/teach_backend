# Intermediate 14: Intermediate capstone: a reliable single-service backend

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 180 minutes. **Outcome:** Demonstrate correct persistent booking, authorization, caching, tests, and a local container stack.

## Prerequisites

Complete [CI checks, structured logs, and measured performance](../13-ci-observability-and-performance/README.md), including its next-class preparation.

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

### 1. This milestone deepens the same project rather than rewarding the number of endpoints.

The hardest checks are simultaneous reservation, ownership enforcement, and behavior when dependencies fail.

### 2. A demonstration is reproducible evidence.

Keep scripts and expected outcomes together, reset fixtures safely, and explain why a test targets a real failure mode.

### 3. The 50% milestone means completing the first two tiers’ coverage goals.

Advanced addresses distributed failure, durable notifications, infrastructure, domain routing, TLS, and operational recovery.

## Exit check

- All mandatory intermediate gates pass.
- Evidence includes tests and live failure demonstrations.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
