# Advanced 06: Linux operations, network boundaries, and secrets

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 180 minutes. **Outcome:** Prepare a host with limited access and trace network connectivity between services.

## Prerequisites

Complete [Threat modeling, OTP, and abuse controls](../05-threat-modeling-and-otp/README.md), including its next-class preparation.

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

### 1. A process runs as a user and listens on selected interfaces.

Bind public entrypoints deliberately; databases and Redis should be reachable only by intended applications. Explain loopback, private network, and public interface with a packet path.

### 2. Use SSH keys, least-privilege accounts, and a documented firewall policy.

Keep an existing administrative session open while testing access changes to avoid locking yourself out.

### 3. Secrets need controlled injection, rotation, and restricted access.

Environment variables and Kubernetes Secrets are delivery mechanisms, not automatic protection against every reader or log leak. Keep secrets out of images, source control, and command transcripts.

## Exit check

- Only intended entrypoints are exposed.
- Rotation works without a source-code edit.
- Recovery access is tested before closing the original session.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
