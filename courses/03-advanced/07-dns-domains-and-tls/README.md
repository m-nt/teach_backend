# Advanced 07: DNS, domains, reverse proxies, and TLS

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 180 minutes. **Outcome:** Route a hostname to the API and verify HTTPS with an explicit trust model.

## Prerequisites

Complete [Linux operations, network boundaries, and secrets](../06-linux-networking-and-secrets/README.md), including its next-class preparation.

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

### 1. DNS maps a hostname to records, such as A/AAAA addresses.

DNS resolution, TCP reachability, TLS validation, and HTTP routing are separate layers; troubleshoot them in that order.

### 2. A reverse proxy accepts public traffic and forwards to a private API.

TLS protects traffic and authenticates the server name when the client trusts the issuing authority. “SSL” is common shorthand; teach modern TLS terminology.

### 3. Public automatic certificates require control of the name and a successful validation challenge.

For offline/local practice use Caddy internal TLS and explicitly trust the local root on the lab client. Do not teach disabling certificate verification as the solution.

## Exit check

- Client verifies the certificate without -k.
- The chosen hostname reaches the API.
- Student can isolate a DNS failure from a TLS failure.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
