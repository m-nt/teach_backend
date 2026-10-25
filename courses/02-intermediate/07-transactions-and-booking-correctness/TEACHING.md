# Instructor notes: Transactions and concurrent booking correctness

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. A transaction groups changes into one commit or rollback. Reading availability and inserting later without coordination allows two callers to take the last seat.

2. For this project, lock the session row with SELECT FOR UPDATE, count active bookings, enforce capacity, insert, then commit. Every reservation path must acquire that same lock before deciding. A unique student/session constraint separately prevents duplicate reservations.

3. A transaction’s lifetime should be short. Do not send email while holding the database lock. Return only after commit succeeds, and translate conflicts into the agreed HTTP contract.

## Demonstration script

In two psql sessions reproduce a last-seat race, then repeat with the shared session-row lock. Show the second transaction waits and then sees the committed booking.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- Why is a pre-check outside the transaction insufficient? **Expected reasoning:** Another caller can change availability between the check and write.
- What does the unique pair constraint prevent? **Expected reasoning:** Two bookings for the same student and session.
- Why keep network calls outside the transaction? **Expected reasoning:** Slow or failed calls prolong locks and couple unrelated failures.

## Review the lab

- Exactly one request succeeds and capacity is not exceeded.
- Duplicate student/session requests are rejected.
- No external notification runs inside the lock-holding transaction.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
