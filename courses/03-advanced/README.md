# Advanced

[Course overview](../../README.md) · [Assessment](../../instructor/assessment.md) · [Capstone](CAPSTONE.md)

14 lessons × 180 minutes = 42 contact hours. Homework is additional: 4–6 hours per lesson.

## Entry and exit

Entry: intermediate capstone or equivalent practical assessment. Exit: deploy and operate the defined multi-service lab, demonstrate verified TLS, durable work, abuse controls, Kubernetes operations, and restore/rollback drills.

## Lesson sequence

| Lesson | Outcome |
|---|---|
| [01 — Architecture, requirements, and service boundaries](01-architecture-and-service-boundaries/README.md) | Define a small multi-service system with explicit ownership and failure assumptions. |
| [02 — Standards, contracts, and compatible evolution](02-standards-and-compatible-contracts/README.md) | Version HTTP and event contracts without breaking deployed consumers. |
| [03 — Transactional outbox and durable notification workers](03-durable-outbox-and-workers/README.md) | Commit booking and delivery intent atomically and process work with recoverable claims. |
| [04 — Timeouts, retries, backpressure, and overload](04-resilience-and-backpressure/README.md) | Bound failure amplification and demonstrate recovery from a slow dependency. |
| [05 — Threat modeling, OTP, and abuse controls](05-threat-modeling-and-otp/README.md) | Implement a constrained one-time verification flow and test replay and guessing defenses. |
| [06 — Linux operations, network boundaries, and secrets](06-linux-networking-and-secrets/README.md) | Prepare a host with limited access and trace network connectivity between services. |
| [07 — DNS, domains, reverse proxies, and TLS](07-dns-domains-and-tls/README.md) | Route a hostname to the API and verify HTTPS with an explicit trust model. |
| [08 — Delivery pipelines, release rollback, and safe migrations](08-delivery-and-safe-migrations/README.md) | Deploy a versioned release and roll back application code without corrupting data. |
| [09 — Kubernetes: Pods, Deployments, Services, and configuration](09-kubernetes-foundations/README.md) | Deploy the business API and notification worker to a local Kubernetes cluster. |
| [10 — Kubernetes probes, resources, rollout, and network controls](10-kubernetes-operations/README.md) | Operate a bounded rollout and diagnose unhealthy or unschedulable workloads. |
| [11 — Stateful services, backups, and restore drills](11-backups-state-and-recovery/README.md) | Restore PostgreSQL to an isolated target and verify business-level recovery. |
| [12 — Metrics, traces, SLOs, and incident diagnosis](12-observability-and-slos/README.md) | Diagnose a staged incident using user-facing objectives and correlated telemetry. |
| [13 — Infrastructure capstone integration and rehearsal](13-infrastructure-capstone-workshop/README.md) | Assemble the multi-service deployment and rehearse its failure and recovery checks. |
| [14 — Final demonstration and operational handoff](14-final-demo-and-operational-handoff/README.md) | Demonstrate a small operated infrastructure and defend its guarantees and limitations. |

## Teaching practice

Teach in order. Each folder has a student lesson, instructor notes, a lab, a quiz with a separated answer key, and homework with next-class preparation. Use the exit check before adding complexity. If most learners cannot complete it, use a clinic and rebaseline the schedule rather than skip the prerequisite.

Core labs fit class time with instructor guidance. Homework completes and practises the slice; extension work is optional. Later-tier labs extend the student capstone and deliberately do not provide a complete copy-paste solution. Use the fallback demo and answer guidance in each lesson to unblock discussion.
