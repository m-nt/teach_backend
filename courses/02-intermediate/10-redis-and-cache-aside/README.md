# Intermediate 10: Redis and cache-aside behavior

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 180 minutes. **Outcome:** Cache public session reads with explicit expiry, invalidation, and a database fallback.

## Prerequisites

Complete [Password storage, sessions, and resource authorization](../09-authentication-and-authorization/README.md), including its next-class preparation.

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

### 1. A cache holds a derived copy of data.

In cache-aside, read the cache, load the database on a miss, and populate with a bounded TTL. The database remains the source of truth.

### 2. Invalidation matters more than the happy-path hit.

Delete or refresh affected keys after a successful write. A stale capacity display must never authorize a booking; the transactional rule still checks the database.

### 3. Cache failure should not silently change correctness.

For public catalog reads, fall back to the database. Scope keys correctly and do not cache private responses under a shared public key.

## Exit check

- An update becomes visible after invalidation.
- Booking authority never depends on the cache.
- Redis outage does not break the public read.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
