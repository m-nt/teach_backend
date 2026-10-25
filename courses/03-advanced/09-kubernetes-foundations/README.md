# Advanced 09: Kubernetes: Pods, Deployments, Services, and configuration

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 180 minutes. **Outcome:** Deploy the business API and notification worker to a local Kubernetes cluster.

## Prerequisites

Complete [Delivery pipelines, release rollback, and safe migrations](../08-delivery-and-safe-migrations/README.md), including its next-class preparation.

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

### 1. A Pod is the scheduling unit; a Deployment manages replica rollout and replacement; a Service supplies stable discovery for selected Pods.

Containers and Pods are not interchangeable concepts.

### 2. A namespace groups resources, ConfigMaps carry non-secret configuration, and Secrets carry sensitive values with access controls that must be configured.

Resource requests help scheduling and limits constrain consumption.

### 3. Use a single-node training cluster to learn the control loop.

It does not demonstrate high availability. Keep PostgreSQL external to the first stateless deployment or use the explicitly disposable lab database track.

## Exit check

- API and worker run as separate Deployments.
- Service reaches replacement Pods.
- No real secret is committed to the manifests.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
