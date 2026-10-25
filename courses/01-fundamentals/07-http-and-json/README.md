# Fundamentals 07: HTTP, URLs, JSON, and API contracts

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 150 minutes. **Outcome:** Read a request and response and design a tiny booking API contract.

## Prerequisites

Complete [Git, debugging, and a small project structure](../06-git-debugging-and-project-shape/README.md), including its next-class preparation.

## Class plan

| Minutes | Activity |
|---|---|
| 00–10 | Retrieval practice and setup check |
| 10–35 | Explain concepts with the whiteboard example |
| 35–55 | Instructor live demonstration |
| 55–65 | Break |
| 65–110 | Guided lab and pair explanation |
| 110–130 | Independent lab change and review |
| 130–140 | Quiz and correction |
| 140–150 | Exit check, homework, next-class preparation |

## Lesson

### 1. HTTP connects a client request with a server response.

A URL contains a scheme, host, optional port, path, and possibly a query string. A method expresses intent; headers carry metadata; the body carries data.

### 2. JSON represents objects, arrays, strings, numbers, booleans, and null.

It is a wire format, not a database. Distinguish a valid JSON document from a document with valid business values.

### 3. Agree on a contract before implementing handlers: GET /sessions lists classes; POST /bookings requests a booking.

Use 200 for a successful read, 201 for creation, 400 for malformed input, 404 for absent resources, and 409 for a full class conflict. These are this course’s contract choices.

## Exit check

- Student identifies method, path, status, headers, and body.
- Contract contains success and failure examples.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
