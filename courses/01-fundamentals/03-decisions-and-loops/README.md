# Fundamentals 03: Decisions, loops, and tracing

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 150 minutes. **Outcome:** Implement booking decisions with if, switch, and bounded loops.

## Prerequisites

Complete [Values, variables, types, and scope](../02-values-variables-and-scope/README.md), including its next-class preparation.

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

### 1. A condition evaluates to true or false.

A booking is allowed only when the class is open and remaining seats are positive. Translate the sentence into small conditions before combining them with &&.

### 2. A loop repeats work.

Trace one iteration at a time on paper with columns for attempt and remaining seats. The loop condition is checked before each iteration; updating the wrong variable can prevent termination.

### 3. Boundary values expose mistakes: zero seats, exactly one seat, and a closed class with seats available.

switch is useful for named states; it does not replace all if statements.

## Exit check

- Five attempts accept exactly three and reject two.
- A closed class accepts no requests.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
