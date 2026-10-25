# Advanced 13: Infrastructure capstone integration and rehearsal

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 180 minutes. **Outcome:** Assemble the multi-service deployment and rehearse its failure and recovery checks.

## Prerequisites

Complete [Metrics, traces, SLOs, and incident diagnosis](../12-observability-and-slos/README.md), including its next-class preparation.

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

### 1. The capstone integrates a business API, notification worker, PostgreSQL, Redis, sandbox mail, hostname routing, and verified TLS.

Kubernetes runs the stateless workloads; document where stateful dependencies live and why.

### 2. Keep a deployment inventory with image versions, configuration names, ports, volumes, migrations, and recovery commands.

A peer should be able to operate the system without your memory.

### 3. Use staged failures to test guarantees: replay a booking, race OTP verification, stop Redis, crash the worker, replace a Pod, and restore a backup.

Fix core correctness before adding optional services.

## Exit check

- All required services and ownership boundaries are documented.
- Failures have observable, controlled outcomes.
- Remaining limitations are explicit.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
