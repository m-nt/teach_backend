# Fundamentals 06: Git, debugging, and a small project structure

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 150 minutes. **Outcome:** Commit a working program and navigate a minimal Go project without premature layers.

## Prerequisites

Complete [Structs, slices, maps, and business entities](../05-structs-slices-and-entities/README.md), including its next-class preparation.

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

### 1. Git records snapshots of tracked files.

The working tree, staging area, and commit are different states. git status explains what is about to be included. A commit is local; a push sends it to a remote.

### 2. A Go module has a go.mod and one or more packages.

Start with a single package while the program is small. Explain main.go, a rule file, and a test file; do not force a directory per concept.

### 3. Debug by reproducing, narrowing, changing one thing, and checking again.

Read errors and use temporary prints or an editor breakpoint. Never commit credentials, environment files, or database dumps.

## Exit check

- git log shows two meaningful commits.
- go run . works with multiple source files; no secrets are tracked.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
