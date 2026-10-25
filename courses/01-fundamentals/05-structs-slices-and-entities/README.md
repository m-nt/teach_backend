# Fundamentals 05: Structs, slices, maps, and business entities

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 150 minutes. **Outcome:** Model a small class catalog using structs, slices, and stable IDs.

## Prerequisites

Complete [Functions, parameters, and explicit errors](../04-functions-and-errors/README.md), including its next-class preparation.

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

### 1. A struct groups related fields, such as ID, Title, and Capacity.

Its fields describe one entity. Exported names start with an uppercase letter; package visibility is introduced here and revisited later.

### 2. A slice holds a sequence of values; append adds values and range visits them.

A map looks up a value by key. Use a slice for a small ordered catalog and demonstrate a map for lookup, without requiring both everywhere.

### 3. Two classes can have the same title.

Their IDs must differ. A Course describes the subject; a ClassSession is a particular offering at a time. A Booking connects a student to a session. Model that distinction before writing database tables.

## Exit check

- Two sessions with the same title remain distinguishable.
- An absent ID is handled without crashing.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
