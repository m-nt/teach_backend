# Instructor notes: Table-driven tests, HTTP tests, and integration boundaries

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. Table-driven tests describe cases using inputs and expected outcomes. Name cases after behavior, especially boundaries and failures. An assertion should fail if the requirement is broken.

2. httptest records responses without a public server. Test status, response shape, and the absence of unintended state changes. Avoid asserting incidental whitespace or private implementation details.

3. Unit tests isolate rules; integration tests verify the real database boundary. Keep database test fixtures isolated and cleanup explicit so test order does not matter.

## Demonstration script

Create a table for valid, full, invalid, and missing-session requests. Use httptest.NewRecorder and NewRequest for malformed JSON. Show one intentionally broken handler failing.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- What should a test assertion describe? **Expected reasoning:** Observable required behavior.
- What does httptest avoid needing? **Expected reasoning:** A separately deployed HTTP server for handler tests.
- Why isolate database fixtures? **Expected reasoning:** To avoid order dependence and interference between tests or users.

## Review the lab

- Tests fail for a deliberate wrong status or extra booking.
- Integration fixture does not modify classmates’ data.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
