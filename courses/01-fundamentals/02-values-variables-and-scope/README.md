# Fundamentals 02: Values, variables, types, and scope

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 150 minutes. **Outcome:** Use strings, integers, booleans, assignments, and local scope to describe a class.

## Prerequisites

Complete [What a backend does; your first Go program](../01-backend-and-first-program/README.md), including its next-class preparation.

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

### 1. A type describes the values an expression can hold.

Use string for a title, int for seat counts, and bool for an open/closed flag. Start with concrete values rather than explaining every numeric type.

### 2. A variable is a named value that can change.

:= declares a local variable; = assigns an existing one. Constants name values that do not change. A variable declared inside a block is unavailable outside that block.

### 3. A compiler error is feedback, not a failed student.

Read the first error, find the named line, make one change, and run again. fmt.Println prints values; fmt.Printf adds formatting, such as %s and %d.

## Exit check

- Capacity 12 and booked 5 prints remaining 7.
- Student explains why assigning a string to booked fails.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
