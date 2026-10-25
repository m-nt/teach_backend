# Fundamentals 04: Functions, parameters, and explicit errors

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 150 minutes. **Outcome:** Extract a booking rule into a function and handle its error at the caller.

## Prerequisites

Complete [Decisions, loops, and tracing](../03-decisions-and-loops/README.md), including its next-class preparation.

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

### 1. A function names a task.

Parameters are inputs and return values are outputs. Start with remaining(capacity, booked int) int, then extract a rule that can fail.

### 2. Go commonly returns a value and an error.

nil means no error. A caller should check the error before using a result that may be invalid. Returning an error lets the caller choose whether to print a message or later send an HTTP response.

### 3. Keep business rules separate from presentation.

A booking function should return a result rather than printing every decision. That same function can serve a CLI now and an HTTP handler later.

## Exit check

- A full class returns an error and keeps the original state.
- A valid request increments booked once.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
