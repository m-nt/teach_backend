# Advanced 08: Delivery pipelines, release rollback, and safe migrations

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 180 minutes. **Outcome:** Deploy a versioned release and roll back application code without corrupting data.

## Prerequisites

Complete [DNS, domains, reverse proxies, and TLS](../07-dns-domains-and-tls/README.md), including its next-class preparation.

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

### 1. Build once and promote the same immutable artifact.

Record source revision and image digest; floating tags make incident reconstruction harder. Run quality checks before promotion.

### 2. Use expand/migrate/contract for schema changes that overlap old and new application versions.

A rollback of code does not automatically reverse data changes. Backups do not replace a compatibility plan.

### 3. Deployment evidence includes health, a synthetic transaction, error metrics, and an explicit rollback trigger.

Migration jobs should run once under controlled ownership, not independently in every replica.

## Exit check

- Release evidence names a digest and source revision.
- Rollback completes without losing bookings.
- Migration ownership is explicit.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
