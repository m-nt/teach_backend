# Instructor notes: Connecting Go to PostgreSQL

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. A database driver speaks the database protocol; database/sql offers a connection pool API. Open the pool once for the application and close it at shutdown. Do not create a new pool per request.

2. Use placeholders and separate arguments for user input. Building SQL by concatenating input can change the meaning of the query. Distinguish no matching row from an unavailable database.

3. A migration is a versioned schema change. For today use the supplied schema and seeded database. The example lists and creates students; implementing transaction-safe booking waits for intermediate. This keeps a first database lesson small enough to understand.

## Demonstration script

Run examples/05-db-api. Trace GET /students and POST /students, including the $1 placeholder and Scan calls. Insert a student, restart Go, and read the student back.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- What is a connection pool for? **Expected reasoning:** Reusing and managing database connections across operations.
- Why use placeholders? **Expected reasoning:** To bind input as data instead of executable SQL syntax.
- Does a schema constraint replace all application validation? **Expected reasoning:** No; both enforce different parts of correctness and usable error handling.

## Review the lab

- New student survives a Go process restart.
- The query uses arguments, not SQL string concatenation.
- Database failure produces an error response rather than a crash.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
