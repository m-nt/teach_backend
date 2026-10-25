# Intermediate 12: Containerizing Go, configuration, and graceful shutdown

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 180 minutes. **Outcome:** Build a Go application image and run a configured local stack with graceful shutdown.

## Prerequisites

Complete [Email adapters and bounded background work](../11-email-and-background-work/README.md), including its next-class preparation.

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

### 1. A multi-stage build compiles Go in a build stage and copies the binary into a smaller runtime stage.

Choose a runtime that meets the program’s needs, such as CA certificates for outbound TLS.

### 2. Configuration should be validated at startup.

Missing required secrets or malformed addresses should fail clearly. Containers on a Compose network address dependencies by service name, not host localhost.

### 3. Graceful shutdown stops accepting new work, lets bounded in-flight work finish, and closes resources.

Readiness means able to serve; liveness means the process should remain running. Their checks need different failure behavior.

## Exit check

- Image starts without local source files.
- API connects using Compose service names.
- Shutdown finishes or times out predictably.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
