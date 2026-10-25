# Intermediate

[Course overview](../../README.md) · [Assessment](../../instructor/assessment.md) · [Capstone](CAPSTONE.md)

14 lessons × 180 minutes = 42 contact hours. Homework is additional: 4–6 hours per lesson.

## Entry and exit

Entry: fundamentals capstone or equivalent placement demonstration. Exit: a tested, authenticated, transaction-safe single-service backend with cache, sandbox email, containers, and quality gates.

## Lesson sequence

| Lesson | Outcome |
|---|---|
| [01 — Go value semantics, pointers, and methods](01-go-values-and-pointers/README.md) | Predict copying and mutation across structs, slices, pointers, and methods. |
| [02 — Packages, small interfaces, and error contracts](02-packages-interfaces-and-errors/README.md) | Separate HTTP, booking behavior, and storage using a small consumer-owned interface. |
| [03 — Table-driven tests, HTTP tests, and integration boundaries](03-tests-and-http-testing/README.md) | Test business behavior and handlers with controlled fixtures and useful assertions. |
| [04 — Context, timeouts, defer, and resource lifetimes](04-context-and-resource-lifetimes/README.md) | Propagate cancellation and close resources without leaking work. |
| [05 — Goroutines, channels, mutexes, and bounded work](05-concurrency-and-races/README.md) | Implement bounded concurrent work and distinguish memory races from business races. |
| [06 — Schema migrations, indexes, and query plans](06-schema-migrations-and-query-plans/README.md) | Evolve a schema safely and measure an index against a real query shape. |
| [07 — Transactions and concurrent booking correctness](07-transactions-and-booking-correctness/README.md) | Prevent overbooking under simultaneous requests using a database transaction. |
| [08 — API contracts, pagination, and middleware](08-api-contracts-and-middleware/README.md) | Publish a consistent HTTP contract with pagination, validation, and request context. |
| [09 — Password storage, sessions, and resource authorization](09-authentication-and-authorization/README.md) | Implement reviewed password hashing and session-based access with ownership checks. |
| [10 — Redis and cache-aside behavior](10-redis-and-cache-aside/README.md) | Cache public session reads with explicit expiry, invalidation, and a database fallback. |
| [11 — Email adapters and bounded background work](11-email-and-background-work/README.md) | Send sandbox email through an adapter and explain the limitations of in-memory jobs. |
| [12 — Containerizing Go, configuration, and graceful shutdown](12-containers-config-and-shutdown/README.md) | Build a Go application image and run a configured local stack with graceful shutdown. |
| [13 — CI checks, structured logs, and measured performance](13-ci-observability-and-performance/README.md) | Create a reproducible quality gate and diagnose one request using logs and measurements. |
| [14 — Intermediate capstone: a reliable single-service backend](14-intermediate-capstone/README.md) | Demonstrate correct persistent booking, authorization, caching, tests, and a local container stack. |

## Teaching practice

Teach in order. Each folder has a student lesson, instructor notes, a lab, a quiz with a separated answer key, and homework with next-class preparation. Use the exit check before adding complexity. If most learners cannot complete it, use a clinic and rebaseline the schedule rather than skip the prerequisite.

Core labs fit class time with instructor guidance. Homework completes and practises the slice; extension work is optional. Later-tier labs extend the student capstone and deliberately do not provide a complete copy-paste solution. Use the fallback demo and answer guidance in each lesson to unblock discussion.
