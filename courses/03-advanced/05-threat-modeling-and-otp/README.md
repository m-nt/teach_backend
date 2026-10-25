# Advanced 05: Threat modeling, OTP, and abuse controls

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 180 minutes. **Outcome:** Implement a constrained one-time verification flow and test replay and guessing defenses.

## Prerequisites

Complete [Timeouts, retries, backpressure, and overload](../04-resilience-and-backpressure/README.md), including its next-class preparation.

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

### 1. An OTP proves temporary access to a delivery channel; it is not automatically strong multi-factor authentication.

Bind a challenge to purpose and account or destination. Generate codes using crypto/rand, not math/rand.

### 2. Short codes have low entropy.

Store a keyed digest using a secret held outside the challenge store, expire quickly, limit attempts and issuance, and atomically verify-and-consume. Never log codes, tokens, or full destinations. Resending must not create unlimited concurrent valid codes.

### 3. Generic responses reduce account enumeration.

Rate limits should combine relevant account/destination and network signals, with documented behavior when the limiter store fails. Do not use the public-catalog cache fallback policy for authentication controls.

## Exit check

- Only one concurrent verification succeeds.
- Wrong attempts and issuance are bounded.
- No code or session token appears in logs.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
