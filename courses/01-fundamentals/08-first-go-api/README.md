# Fundamentals 08: Building a small Go HTTP API

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 150 minutes. **Outcome:** Connect handlers, JSON decoding, validation, and booking rules.

## Prerequisites

Complete [HTTP, URLs, JSON, and API contracts](../07-http-and-json/README.md), including its next-class preparation.

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

### 1. net/http calls a handler for a request.

The handler translates HTTP input into Go values, invokes a rule, and translates the result into HTTP output. Keep transport details out of the core booking rule.

### 2. Decode JSON into a request struct, reject invalid values, and set the content type and status before writing the response.

A missing error check can turn a bad request into a misleading success.

### 3. The example keeps data in process memory, so restarting loses bookings.

HTTP handlers may execute concurrently; the supplied example locks its state. Students trace that lock today; concurrency mechanics belong to intermediate.

## Exit check

- /about returns JSON with HTTP 200.
- Invalid names produce 400 and do not consume seats.
- Restart behavior is explained correctly.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
