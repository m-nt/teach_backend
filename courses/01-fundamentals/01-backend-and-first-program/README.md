# Fundamentals 01: What a backend does; your first Go program

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 150 minutes. **Outcome:** Explain a request journey and edit, run, and deliberately break a Go program.

## Prerequisites

No coding prerequisite. Complete the supported setup or use an instructor-prepared machine.

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

### 1. A program is a sequence of instructions executed by a computer.

Source code is the text we edit; Go compiles that text into a program. A terminal is a way to give commands, and its working directory decides which files a command sees. Demonstrate folders before introducing syntax.

### 2. A customer asks to book a class.

The browser sends a request; the backend checks rules; a database remembers the booking; a response tells the browser what happened. Draw these as four boxes. A database stores facts but does not decide every business rule.

### 3. A business entity is something the business identifies over time: a Student, Course, ClassSession, or Booking.

An ID identifies an entity even when its name changes. Business identity here means domain identity, not company branding. Today students only name entities; tables come later.

## Exit check

- A greeting printed from the student’s own file.
- A student can recover from one missing quote and explain the four boxes.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
