# Advanced 12: Metrics, traces, SLOs, and incident diagnosis

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 180 minutes. **Outcome:** Diagnose a staged incident using user-facing objectives and correlated telemetry.

## Prerequisites

Complete [Stateful services, backups, and restore drills](../11-backups-state-and-recovery/README.md), including its next-class preparation.

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

### 1. An SLI is a measured indicator; an SLO is its target over a stated window.

Define successful booking requests and acceptable latency precisely, including exclusions. A single successful probe is not a service-level history.

### 2. Metrics reveal rates, errors, durations, saturation, and queue lag.

Keep labels bounded; user IDs and arbitrary URLs create excessive cardinality. Logs explain events; traces connect work across API, storage, and worker boundaries.

### 3. Alert on actionable user impact and clear ownership.

Use a runbook to identify whether the incident is in DNS, proxy, API, database, cache, queue, or mail delivery. Measure before changing multiple things.

## Exit check

- Telemetry distinguishes API success from delayed email delivery.
- No secret or user ID is used as an unbounded metric label.
- Incident diagnosis cites observations.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
