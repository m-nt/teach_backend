# Intermediate capstone: reliable booking API

Extend the beginner project into one well-structured business service. The source of truth is PostgreSQL; Redis caches public reads; a sandbox email adapter demonstrates notifications.

## Required behavior and evidence

| Requirement | Demonstration |
|---|---|
| Persistent reservation | Create booking, restart API, read it back |
| No overbooking | Ten concurrent requests from distinct synthetic users for one remaining seat; exactly one succeeds |
| No duplicate reservation | Same student/session pair rejected by application and database constraint |
| Authentication | Login with a password-specific hash implementation; session expiry and logout work |
| Authorization | Student A cannot read or cancel student B’s booking; staff permission is explicit |
| Browser security | Cookie settings and CSRF defense tested for state-changing routes |
| HTTP contract | Documented methods, fields, statuses, pagination bounds, stable errors |
| Cache | Hit/miss/expiry/invalidation and Redis-down database fallback |
| Email | Fake-sender test and message delivered only to Mailpit; in-memory loss window documented |
| Resource lifetimes | Request cancellation, bounded workers, graceful shutdown |
| Quality | Rule/handler/integration tests, race check, vet/build, reviewed migration |
| Delivery | API image and Compose runbook reproduced by a peer |

## Recommended order

Migrate the schema; implement the transaction-safe rule; connect transport; add tests; add authentication/ownership; add cache and email; package and observe. Keep scope fixed after lesson 13.

## Submission

Include source revision, schema/migrations, API contract, automated test results, workload setup for the last-seat test, container commands, redacted logs, and a short dependency diagram. Record test environment and actual versions. The supplied infrastructure starts dependencies only; your application image and wiring are part of the project.

Use the [assessment rubric](../../instructor/assessment.md). The mandatory capacity and authorization gates cannot be replaced with a higher quiz score. Advanced begins only when these foundations work.
