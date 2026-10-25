# Intermediate 06: Schema migrations, indexes, and query plans

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 180 minutes. **Outcome:** Evolve a schema safely and measure an index against a real query shape.

## Prerequisites

Complete [Goroutines, channels, mutexes, and bounded work](../05-concurrency-and-races/README.md), including its next-class preparation.

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

### 1. A migration records an ordered schema change.

Separate schema changes from ad hoc edits, and test both an empty database and an upgrade with existing data. Rolling deployments may run old and new application versions together.

### 2. An index trades storage and write work for faster access patterns.

Choose indexes from WHERE, JOIN, and ORDER BY requirements. A tiny table may reasonably use a sequential scan.

### 3. Use EXPLAIN and EXPLAIN ANALYZE to inspect plans; ANALYZE executes the statement.

Use synthetic read queries for this lab and compare results at a meaningful data size.

## Exit check

- Existing booking data survives the upgrade.
- Plan comparison includes row count and query text.
- Rollback data-loss consequences are stated.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
