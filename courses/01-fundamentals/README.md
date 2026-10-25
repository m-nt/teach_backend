# Fundamentals

[Course overview](../../README.md) · [Assessment](../../instructor/assessment.md) · [Capstone](CAPSTONE.md)

14 lessons × 150 minutes = 35 contact hours. Homework is additional: 2–4 hours per lesson.

## Entry and exit

Entry: no programming experience. Learners need basic reading, typing, and file navigation, with supported setup time. Exit: run and explain a small API, model data, write basic Go/SQL, use Git and Docker, and demonstrate persistent registration.

## Lesson sequence

| Lesson | Outcome |
|---|---|
| [01 — What a backend does; your first Go program](01-backend-and-first-program/README.md) | Explain a request journey and edit, run, and deliberately break a Go program. |
| [02 — Values, variables, types, and scope](02-values-variables-and-scope/README.md) | Use strings, integers, booleans, assignments, and local scope to describe a class. |
| [03 — Decisions, loops, and tracing](03-decisions-and-loops/README.md) | Implement booking decisions with if, switch, and bounded loops. |
| [04 — Functions, parameters, and explicit errors](04-functions-and-errors/README.md) | Extract a booking rule into a function and handle its error at the caller. |
| [05 — Structs, slices, maps, and business entities](05-structs-slices-and-entities/README.md) | Model a small class catalog using structs, slices, and stable IDs. |
| [06 — Git, debugging, and a small project structure](06-git-debugging-and-project-shape/README.md) | Commit a working program and navigate a minimal Go project without premature layers. |
| [07 — HTTP, URLs, JSON, and API contracts](07-http-and-json/README.md) | Read a request and response and design a tiny booking API contract. |
| [08 — Building a small Go HTTP API](08-first-go-api/README.md) | Connect handlers, JSON decoding, validation, and booking rules. |
| [09 — Docker, images, containers, ports, and volumes](09-docker-and-local-services/README.md) | Start a local PostgreSQL container and distinguish process lifecycle from stored data. |
| [10 — Relational data modeling and basic SQL](10-sql-and-data-modeling/README.md) | Create related tables and use basic SQL to store and query bookings. |
| [11 — Connecting Go to PostgreSQL](11-go-and-postgres/README.md) | Trace and run a database-backed endpoint with parameterized SQL and explicit error handling. |
| [12 — Testing, configuration, and security basics](12-testing-config-and-security-basics/README.md) | Write a simple Go test and recognize common unsafe backend behavior. |
| [13 — Fundamentals capstone workshop](13-capstone-workshop/README.md) | Assemble and explain a small backend with a clear domain model and working persistence. |
| [14 — Capstone demonstrations and fundamentals review](14-demonstration-and-next-steps/README.md) | Demonstrate the beginner backend and identify readiness gaps before intermediate. |

## Teaching practice

Teach in order. Each folder has a student lesson, instructor notes, a lab, a quiz with a separated answer key, and homework with next-class preparation. Use the exit check before adding complexity. If most learners cannot complete it, use a clinic and rebaseline the schedule rather than skip the prerequisite.

Core labs fit class time with instructor guidance. Homework completes and practises the slice; extension work is optional. Later-tier labs extend the student capstone and deliberately do not provide a complete copy-paste solution. Use the fallback demo and answer guidance in each lesson to unblock discussion.
