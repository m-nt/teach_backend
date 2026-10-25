# Intermediate 01: Go value semantics, pointers, and methods

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 180 minutes. **Outcome:** Predict copying and mutation across structs, slices, pointers, and methods.

## Prerequisites

Pass this tier’s entry gate and bring the prior capstone.

## Class plan

| Minutes | Activity |
|---|---|
| 00–15 | Retrieval practice and prerequisite check |
| 15–45 | Concepts, design discussion, and whiteboard |
| 45–70 | Live demonstration and failure case |
| 70–80 | Break |
| 80–135 | Guided implementation lab |
| 135–160 | Failure tests and peer review |
| 160–170 | Quiz and correction |
| 170–180 | Exit check and next-class preparation |

## Lesson

### 1. Assigning a struct copies its fields.

A pointer refers to an existing value, so changing through the pointer can affect the caller’s value. Explain ownership using two boxes and an arrow before showing syntax.

### 2. A slice value describes a backing array; copying the slice does not necessarily copy its elements.

append may allocate another array. Code should not depend on accidental spare capacity.

### 3. Methods associate behavior with a type.

Choose pointer receivers when mutation or consistent receiver behavior requires them. A pointer is not automatically faster or safer.

## Exit check

- Prediction matches observed struct and slice behavior.
- Tests prove Reserve mutates only the intended value.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
