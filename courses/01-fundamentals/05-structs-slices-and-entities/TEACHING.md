# Instructor notes: Structs, slices, maps, and business entities

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. A struct groups related fields, such as ID, Title, and Capacity. Its fields describe one entity. Exported names start with an uppercase letter; package visibility is introduced here and revisited later.

2. A slice holds a sequence of values; append adds values and range visits them. A map looks up a value by key. Use a slice for a small ordered catalog and demonstrate a map for lookup, without requiring both everywhere.

3. Two classes can have the same title. Their IDs must differ. A Course describes the subject; a ClassSession is a particular offering at a time. A Booking connects a student to a session. Model that distinction before writing database tables.

## Demonstration script

Walk through examples/02-booking-cli. Rename a session without changing its ID. Show a lookup that succeeds and one that fails.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- Why use IDs instead of titles as identity? **Expected reasoning:** Titles can change or repeat.
- What does a slice represent here? **Expected reasoning:** An ordered collection of sessions.
- Which entity connects a student to a scheduled class? **Expected reasoning:** Booking.

## Review the lab

- Two sessions with the same title remain distinguishable.
- An absent ID is handled without crashing.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
