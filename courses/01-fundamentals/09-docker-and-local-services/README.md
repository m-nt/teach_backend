# Fundamentals 09: Docker, images, containers, ports, and volumes

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 150 minutes. **Outcome:** Start a local PostgreSQL container and distinguish process lifecycle from stored data.

## Prerequisites

Complete [Building a small Go HTTP API](../08-first-go-api/README.md), including its next-class preparation.

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

### 1. An image packages a filesystem and launch configuration; a container is a running instance.

Docker is a tool for running dependencies consistently, not a replacement for understanding them.

### 2. A published port connects a host port to a container port.

localhost inside a container refers to that container. Compose service names identify other services on its network. Keep the training database bound to the host loopback interface.

### 3. A volume stores data separately from a container lifecycle.

Stopping or replacing a container is different from deleting its volume. The local Compose file uses development credentials and must stay local.

## Exit check

- SELECT 1 returns one row.
- Student identifies host port 5432 and container port 5432.
- Student explains where database files persist.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
