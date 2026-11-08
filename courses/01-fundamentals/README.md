# Fundamentals

[Course overview](../../README.md) · [Assessment](../../instructor/assessment.md) · [Capstone](CAPSTONE.md)

18 lessons = **17 × 120 minutes + 1 × 60-minute demonstration = 2,100 minutes = 35 contact hours**. Homework is additional: 2–4 hours per lesson. The schedule spans about 12 weeks when meetings are distributed across the calendar.

This tier starts with no programming knowledge. It teaches the request/response journey, Go basics, booking rules, entities, Git, networking, TCP, HTTP, APIs, Docker, SQL, persistence, tests, and basic security. The sequence deliberately builds each concept before asking students to connect it to the next layer.

## Entry and exit

Entry: no programming experience. Learners need basic reading, typing, and file navigation, with supported setup time. Exit: run and explain a small API, model data, write basic Go/SQL, use Git and Docker, and demonstrate persistent registration. Completion represents introductory breadth—the first 20% of this program’s planned competency journey—not production readiness.

## Lesson sequence

| Lesson | Topic | Practical output |
|---|---|---|
| [01 — Backend and the request/response cycle](01-backend-and-request-response/README.md) | Backend responsibilities and request journey | Draw a request from client to server |
| [02 — First Go program, variables, and data types](02-first-go-program-variables-and-types/README.md) | Program structure and typed values | Build a small program |
| [03 — Conditions, loops, and control flow](03-conditions-loops-and-control-flow/README.md) | Branches and bounded repetition | Implement booking-system rules |
| [04 — Functions and error handling](04-functions-and-error-handling/README.md) | Contracts and errors as values | Separate business rules |
| [05 — Structs, slices, maps, and data modeling](05-structs-slices-maps-and-data-modeling/README.md) | Entities and collections | Build initial booking models |
| [06 — Git, debugging, and project structure](06-git-debugging-and-project-structure/README.md) | Versioning and diagnosis | Version and debug the project |
| [07 — Network, client and server, IP and port](07-network-client-server-ip-and-port/README.md) | Host/service addresses | Observe client-server communication |
| [08 — TCP, connections, and byte streams](08-tcp-connections-and-byte-streams/README.md) | Connections and message boundaries | Build a simple TCP server |
| [09 — HTTP structure and manual request parsing](09-http-structure-and-manual-request-parsing/README.md) | Method, target, headers, body | Read a raw request and produce a response |
| [10 — HTTP in Go and net/http](10-http-in-go-and-net-http/README.md) | Handlers and the HTTP server | Build the first HTTP server |
| [11 — URLs, methods, headers, status codes, and JSON](11-urls-methods-headers-status-codes-and-json/README.md) | API contract fundamentals | Implement several endpoints |
| [12 — Routing, validation, and API error handling](12-routing-validation-and-api-errors/README.md) | Request dispatch and controlled failures | Build a booking API |
| [13 — Docker and running services](13-docker-and-running-services/README.md) | Images, containers, ports, volumes | Run PostgreSQL in a container |
| [14 — Relational modeling and SQL](14-relational-modeling-and-sql/README.md) | Tables, keys, constraints, queries | Design tables and run queries |
| [15 — Connecting Go to PostgreSQL](15-connecting-go-to-postgresql/README.md) | Pools and parameterized SQL | Persist and retrieve through an API |
| [16 — Testing, configuration, and basic security](16-testing-configuration-and-basic-security/README.md) | Tests, environment settings, common risks | Test rules and identify security issues |
| [17 — Final project](17-final-project/README.md) | Integrated beginner backend | Build a small database-backed backend |
| [18 — Demonstration and review](18-demonstration-and-review/README.md) | Evidence and readiness | Demonstrate the backend and plan bridge work |

## Teaching practice

Teach in order. Each folder contains a student lesson, instructor notes, a lab, a quiz with a separated answer key, and homework with next-class preparation. Use the exit check before adding complexity. If most learners cannot complete it, use a clinic and rebaseline the schedule rather than skip the prerequisite.

The first 17 lessons are 120-minute sessions. Lesson 18 is a 60-minute demonstration and review, so the advertised 35 contact hours remain unchanged. Core labs fit class time with instructor guidance; homework completes and practises the slice. Extension work is optional. Later-tier labs extend the learner’s capstone and deliberately do not provide a complete copy-paste solution.
