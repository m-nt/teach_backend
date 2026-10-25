# Advanced 10: Kubernetes probes, resources, rollout, and network controls

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 180 minutes. **Outcome:** Operate a bounded rollout and diagnose unhealthy or unschedulable workloads.

## Prerequisites

Complete [Kubernetes: Pods, Deployments, Services, and configuration](../09-kubernetes-foundations/README.md), including its next-class preparation.

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

### 1. Readiness controls whether a Pod receives normal Service traffic; liveness can restart a container.

A startup probe allows slow initialization. Dependency outages should not cause every healthy process to restart continuously.

### 2. Set requests and limits based on observation.

Inspect events for scheduling failures and terminated-container reasons for memory exhaustion. Autoscaling needs appropriate metrics and does not fix a database bottleneck.

### 3. NetworkPolicies require a supporting network plugin.

Restrict ingress/egress by the actual dependency graph, then test both allowed and denied flows. RBAC limits API actions; it is different from network access.

## Exit check

- Unready Pods are excluded from Service endpoints.
- The failure is explained using events or logs.
- Network denial is verified or explicitly marked unverified if the plugin lacks support.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
