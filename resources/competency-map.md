# Competency map and roadmap alignment

The 20% / 50% labels are local milestones, not a numerical measure of the backend profession. Use the observable cells below to describe completion. Roadmap categories guide breadth; actual instruction is bounded by hours, prerequisites, and capstone evidence.

| Competency | Fundamentals: introductory breadth | Intermediate: independent small-service work | Advanced: operated small infrastructure |
|---|---|---|---|
| Computer/tool use | Files, terminal, IDE, Go install/run, error reading | Modules, dependency review, repeatable checks | Host access, release artifacts, operational diagnosis |
| Go language | Values, scope, conditions, loops, functions, errors, structs, slices/maps | Pointers, methods, interfaces, packages, contexts, concurrency | Resource limits, cancellation across services, measured optimization |
| Domain and design | Entity IDs, rules, simple project structure | Transport/service/storage boundaries | Service ownership, contracts, architecture decisions |
| Networking/API | URL, HTTP, JSON, status codes | Validation, middleware, pagination, documented contracts | DNS, reverse proxy, TLS, compatible contract evolution |
| Data | Relational modeling, keys, CRUD, joins | Migrations, indexes, query plans, transactions, concurrent correctness | Outbox, state placement, backups and verified restoration |
| Security | Input checks, secret awareness, auth vocabulary | Hashing, sessions, CSRF, ownership tests | Threat modeling, OTP abuse controls, rotation, network/RBAC boundaries |
| Background work | Explain notifications | Sandbox email and bounded workers | Durable intent, retries, duplicate handling, recovery |
| Performance | Recognize dependency cost | Cache-aside, basic benchmarks/profiling | Backpressure, workload evidence, latency/error objectives |
| Delivery | Git, local Docker dependencies | App images, Compose, CI and shutdown | Versioned promotion, safe migrations, Kubernetes rollout/rollback |
| Operations | Read errors and reproduce | Structured logs and health endpoints | Metrics/traces, incidents, SLOs, restore and handoff |

Optional specialization after the course: gRPC/Protobuf, Kafka/RabbitMQ, search, object storage, GraphQL, advanced algorithms, generics in depth, reflection, deep runtime/GC analysis, database replication/failover, cloud IAM, Terraform, and multi-region systems. Add a topic only with additional teaching time and explicit acceptance criteria.
