# Instructor notes: Testing, configuration, and security basics

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. A test calls code with controlled inputs and checks the result. Test a successful booking, a full class, and invalid capacity. Expected results should come from the business rule, not from copying the current output.

2. Configuration changes between environments; code should not need editing to change a port or database address. Read environment variables and keep private values out of Git and logs. An .env file is a convenience, not encryption.

3. Authentication asks who the caller is; authorization asks what that caller may do. The beginner examples have no login and use synthetic data on localhost. Identify password hashing, HTTPS, input limits, parameterized queries, and private logging as concerns without implementing a homemade authentication system.

## Demonstration script

Run go test ./... in examples/02-booking-cli. Break the full-class boundary and watch the test fail. Inspect the HTTP examples for body limits and configuration.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- Why deliberately break a rule while checking a test? **Expected reasoning:** To verify the test can detect that behavior change.
- Is .env automatically secure? **Expected reasoning:** No, it is a plaintext file and needs appropriate handling.
- A logged-in student edits another student’s booking: what is missing? **Expected reasoning:** Resource authorization.

## Review the lab

- The new test fails for the deliberate bug and passes after repair.
- Student distinguishes authentication from authorization.
- No credentials appear in the submitted evidence.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
