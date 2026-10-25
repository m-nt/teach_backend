# Fundamentals 10: Relational data modeling and basic SQL

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 150 minutes. **Outcome:** Create related tables and use basic SQL to store and query bookings.

## Prerequisites

Complete [Docker, images, containers, ports, and volumes](../09-docker-and-local-services/README.md), including its next-class preparation.

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

### 1. A table stores rows of one kind of fact; columns have types.

A primary key identifies a row, a foreign key connects rows, and a unique constraint prevents a duplicate combination.

### 2. Model Student, ClassSession, and Booking separately instead of repeating student names in every booking.

A session has many bookings; a student has many bookings. A unique student/session pair expresses one booking per student per session.

### 3. SELECT reads, INSERT creates, UPDATE changes, and DELETE removes.

WHERE limits the rows affected. Use synthetic data and inspect SELECT results before an update or delete. Parameterized application queries are introduced next lesson.

## Exit check

- Join returns names and session titles for seeded bookings.
- Duplicate student/session pair is rejected.
- Student identifies a primary and a foreign key.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
