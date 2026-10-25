# Instructor notes: Packages, small interfaces, and error contracts

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. Package boundaries should make dependencies understandable. Use cmd/api for startup and internal/booking for behavior only when there is enough code to justify it. Keep the domain independent of HTTP status codes.

2. Interfaces describe required behavior and are satisfied implicitly. Define a small repository interface where it is consumed. A fake makes tests possible, but do not add an interface for every struct.

3. Wrap errors with context and preserve identity using %w. errors.Is identifies a known cause; errors.As extracts a typed error. Decide which errors the HTTP layer maps to 404, 409, or a generic server failure.

## Demonstration script

Extract a BookingService using a repository with one operation. Inject a fake that returns ErrFull. Wrap that error and show errors.Is still succeeds.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- Must a Go type declare that it implements an interface? **Expected reasoning:** No, matching methods satisfy it implicitly.
- Why use %w when wrapping? **Expected reasoning:** It preserves the error chain for inspection.
- Where should HTTP status selection live? **Expected reasoning:** At the transport boundary, not inside the domain rule.

## Review the lab

- Domain package does not import net/http.
- Fake-driven test proves error mapping behavior.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
