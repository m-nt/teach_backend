# Advanced 11: Stateful services, backups, and restore drills

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 180 minutes. **Outcome:** Restore PostgreSQL to an isolated target and verify business-level recovery.

## Prerequisites

Complete [Kubernetes probes, resources, rollout, and network controls](../10-kubernetes-operations/README.md), including its next-class preparation.

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

### 1. A replica is not a backup: accidental deletion can replicate.

Define recovery point objective as acceptable data loss and recovery time objective as acceptable recovery duration. Choose a backup strategy against those requirements.

### 2. A PersistentVolumeClaim asks for storage but does not itself provide backup, off-site durability, or proven failover.

For a small deployment, compare a managed database or dedicated host against operating a database in Kubernetes.

### 3. A backup is useful only if restoration is tested.

Restore to a separate target, apply necessary schema compatibility checks, compare counts and known records, and run a synthetic business flow before considering recovery complete.

## Exit check

- Restored data is validated through SQL and the API.
- Source database remains untouched.
- Observed data loss and recovery duration are compared to targets.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
