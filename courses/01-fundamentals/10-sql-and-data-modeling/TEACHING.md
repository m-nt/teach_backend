# Instructor notes: Relational data modeling and basic SQL

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. A table stores rows of one kind of fact; columns have types. A primary key identifies a row, a foreign key connects rows, and a unique constraint prevents a duplicate combination.

2. Model Student, ClassSession, and Booking separately instead of repeating student names in every booking. A session has many bookings; a student has many bookings. A unique student/session pair expresses one booking per student per session.

3. SELECT reads, INSERT creates, UPDATE changes, and DELETE removes. WHERE limits the rows affected. Use synthetic data and inspect SELECT results before an update or delete. Parameterized application queries are introduced next lesson.

## Demonstration script

Apply examples/04-sql/schema.sql and seed.sql to a fresh course database using the local README. Run queries.sql. Try a duplicate booking and a booking with an absent student ID; read the constraint errors.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- What does a foreign key protect? **Expected reasoning:** A reference must point to an existing referenced row, subject to its configured actions.
- Why not use student name as the primary key? **Expected reasoning:** Names can repeat and change.
- What happens with UPDATE without WHERE? **Expected reasoning:** All rows qualifying for the statement, potentially every row, are updated.

## Review the lab

- Join returns names and session titles for seeded bookings.
- Duplicate student/session pair is rejected.
- Student identifies a primary and a foreign key.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
