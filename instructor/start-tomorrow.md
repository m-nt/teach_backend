# First class: ready-to-teach checklist

[Course overview](../README.md) · [Lesson 01](../courses/01-fundamentals/01-backend-and-first-program/README.md)

For a first meeting on **October 26, 2026**, use this plan; if the cohort starts later, keep the relative sequence and change the date. Do not require prior programming knowledge or a cloud account.

## Tonight: instructor preparation (60–90 minutes)

- [ ] Choose your spoken classroom language; keep code identifiers in English and explain every new term. Print the glossary if internet access is unreliable.
- [ ] Confirm room, projector visibility, power strips, seats, Wi-Fi, and an offline folder with the examples and permitted installers.
- [ ] Use a large editor font. Prepare a terminal in `examples/01-hello` and run `go version`, `go run .`, and `go build .`.
- [ ] Rehearse the intentional missing-quote error and recovery. Keep a known-good copy beside your work.
- [ ] Print or share the lesson, lab, and quiz without displaying the answer key.
- [ ] Send a short message: “Bring a charged laptop if available. No coding experience required. Try the Go/editor setup guide; bring any error. We will arrange pairs for setup problems.”
- [ ] Prepare a paper request-flow exercise and a second working laptop or pairing plan. Students without a working setup still drive a partner’s keyboard and explain the flow.

## Opening words (adapt to your voice)

“By the end of this course, you will understand the basic journey from a button click to a business decision and stored data. Today you will run and change your first Go program. We will build a class-booking backend together. You do not need to understand every symbol immediately. Errors are part of the work; we will learn how to read them.”

## 150-minute script

| Time | What to do | Evidence to collect |
|---|---|---|
| 0–10 | Ask who has used a terminal; check devices and make pairs | Setup checklist, no scored coding pretest |
| 10–20 | Draw student → browser → backend → database → response | Students identify where booking rules belong |
| 20–35 | Explain files, folders, terminal working directory, code, and output | Every learner opens a folder and locates the terminal |
| 35–45 | Run the greeting; explain package, import, main, Println only at a surface level | Students predict the printed line |
| 45–55 | Change text, save, rerun; break a quote and repair | Students locate the compiler’s file and line |
| 55–65 | Break; handle one or two setup issues | Pairing fallback for remaining blockers |
| 65–85 | Learners create their own folder, module, and main.go | Each pair runs a program |
| 85–110 | Swap drivers, change city/course text, rerun | Every learner edits and runs, not just watches |
| 110–130 | Draw a booking journey and name Student, Course, Session, Booking | Each learner explains one entity and a stable ID |
| 130–140 | Three-question quiz, then discuss the answer key | Mark prerequisites needing support |
| 140–150 | Individual exit check and homework/prep briefing | Working greeting or precise setup blocker |

## Board layout

Left: the four-box request journey. Middle: a five-line Go program. Right: new words (program, terminal, backend, database, entity, ID). Add terms only when used. Keep “Course = subject” and “ClassSession = scheduled offering” visible.

## Ready-to-type example

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello from our Go backend class!")
}
```

Create a folder, open a terminal there, run `go mod init example.com/first-class`, save the code as `main.go`, then run `go run .`. Expected output: `Hello from our Go backend class!`.

## Questions to ask aloud

- If I change the text but do not save, which version runs? The saved file.
- If two people share a name, are they the same student? No; identity needs a stable ID.
- If a browser lies about available seats, should the backend believe it? No; the backend checks the rule.
- If the program fails, what do we read first? The first relevant error and its location.

## Closing and follow-up

Assign the actual lesson homework, not a generic “study Go.” Ask students to bring their greeting and practise reopening its folder. Record setup failures and offer a short clinic before lesson 02. If fewer than most learners can run and change a program, spend the clinic on that prerequisite and adjust pacing; do not race ahead into syntax.
