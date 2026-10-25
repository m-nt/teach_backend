# Instructor notes: Transactional outbox and durable notification workers

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. Committing a booking and then separately enqueueing a message has a failure gap. An outbox row written in the same database transaction records delivery intent atomically with the booking.

2. Workers claim bounded batches with a recoverable lease or transaction-safe locking strategy. Persist attempt count, next attempt time, and last safe error. A crashed claim must become eligible again.

3. At-least-once processing can duplicate delivery. Give events stable IDs and make internal consumers idempotent using a unique receipt. External email can still duplicate after an ambiguous send result; document the provider’s actual guarantees.

## Demonstration script

Insert booking and outbox together. Kill the worker after sending but before marking completion, restart it, and explain the duplicate window.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- What failure gap does an outbox close? **Expected reasoning:** The gap between committing business state and separately recording delivery work.
- Does an outbox automatically guarantee exactly-once email? **Expected reasoning:** No.
- Why do claims need recovery? **Expected reasoning:** Workers may crash while holding work.

## Review the lab

- A committed booking always has delivery intent.
- Abandoned claims become processable again.
- Duplicate events do not duplicate internal business effects.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
