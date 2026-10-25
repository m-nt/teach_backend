# Instructor notes: Architecture, requirements, and service boundaries

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. Start from requirements: reserve a seat, notify the student, and operate the service when dependencies fail. A service boundary should represent ownership and a reason to deploy independently, not merely a table.

2. Keep booking and identity in one business API initially; extract a notification worker/service that owns delivery attempts. PostgreSQL owns durable business state, Redis owns disposable cache and bounded ephemeral state, and the mail adapter contacts an SMTP sandbox or provider.

3. A network boundary adds latency and partial failure. Record who owns each table, what callers may assume, and what happens when the notification service is down. A modular monolith remains a valid choice for domains that do not need a split.

## Demonstration script

Draw API → PostgreSQL/outbox and worker → mail sandbox, plus API → Redis. Delete each arrow in turn and ask which user behavior fails and which remains available.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- Does each table need its own service? **Expected reasoning:** No; boundaries follow ownership and operational needs.
- What is partial failure? **Expected reasoning:** One part of the system fails while other parts continue working.
- Why document a delivery guarantee? **Expected reasoning:** Clients and operators need to know what success actually promises.

## Review the lab

- Each table and message has one explicit owner.
- Failure behavior is stated for every dependency.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
