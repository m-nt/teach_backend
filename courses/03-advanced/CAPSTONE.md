# Advanced capstone: operated multi-service lab

Deploy a Go business API and a separate notification worker with PostgreSQL, Redis, sandbox email, OTP verification, hostname routing, and verified TLS. Run the stateless workloads in a local or shared Kubernetes cluster. Document where stateful services live and how they are protected and recovered.

## Mandatory acceptance matrix

| Area | Evidence |
|---|---|
| Architecture | Process/storage ownership diagram, trust boundaries, service-split decision |
| Contracts | Versioned API/event schema; compatible additive-change test |
| Idempotency | Repeated request produces one logical booking; conflicting key reuse rejected |
| Durable work | Booking and outbox commit together; crashed worker claim is recovered |
| Duplicate handling | Replayed event has one internal effect; external-email ambiguity documented |
| Resilience | Retry budget, backoff, cancellation, bounded queue/concurrency, observable overload |
| OTP | Expired/wrong/replayed codes denied; two simultaneous verifies yield one success; issuance and attempts bounded |
| Authorization | Cross-user resource tests still pass after deployment |
| Network / TLS | Chosen hostname routes correctly; certificate verifies without bypass; database/Redis not public |
| Release | Recorded source revision/image digest; safe migration and demonstrated application rollback |
| Kubernetes | API/worker Deployments, Service, probes, resource settings, Pod replacement, least-privilege identity |
| Network policy | Allowed and denied traffic tested on a supporting plugin; arrange a supported lab if the local default does not enforce policies |
| Recovery | Backup restored into a different database; known records and API flow verified; RPO/RTO measured |
| Observability | Request correlation, bounded metrics, queue age, defined SLOs, one incident report |
| Handoff | Peer follows deployment, rollback, restore, and credential-rotation runbooks |

Local internal-CA TLS is an acceptable free fallback when no public domain is available; record how the client trusts the CA. A public-domain track requires a controlled subdomain, appropriate DNS and reachability. Do not claim public DNS/ACME validation if only the local track was exercised.

## Suggested final demonstration (10 minutes per team)

1. Two minutes: architecture, versions, verified HTTPS request.
2. Two minutes: booking plus notification and replay behavior.
3. Two minutes: instructor-selected failure and recovery.
4. Two minutes: OTP/ownership evidence and restored-data evidence.
5. Two minutes: individual explanation questions and operational limitations.

Run long restore/load checks before presentation and retain reproducible commands and timestamped evidence. Adjust the timetable or add review sessions for large cohorts; never pretend a backup restore happened because a diagram shows one.

## Scope and handoff

Use synthetic users and a sandbox mail service. Paid email/SMS integration, highly available database operation, multi-region failover, and public production launch are optional follow-ons requiring their own review. The final result is a small operated training infrastructure. See the [rubric](../../instructor/assessment.md) and [runbook template](../../resources/runbook-template.md).
