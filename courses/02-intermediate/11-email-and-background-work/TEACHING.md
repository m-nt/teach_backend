# Instructor notes: Email adapters and bounded background work

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. Separate notification formatting and delivery from booking rules. An EmailSender interface lets tests record a message without contacting a provider. Use Mailpit locally and synthetic addresses.

2. An in-memory queue can demonstrate asynchronous work but loses queued work on process termination. Return a response that reflects the real delivery guarantee. Advanced adds a transactional outbox for durable delivery intent.

3. Retries need a limit and a distinction between transient and permanent errors. A timeout after sending may mean the receiver got the message, so blind retries can duplicate delivery.

## Demonstration script

Send a booking confirmation to Mailpit, inspect the inbox, then stop the mail service and trace the bounded failure path. Kill a worker with queued messages to reveal the durability gap.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- Is an in-memory queue durable across restarts? **Expected reasoning:** No.
- Does a timeout prove the email was not delivered? **Expected reasoning:** No, delivery may have happened before the timeout.
- Why use a fake sender in tests? **Expected reasoning:** To verify behavior without depending on or contacting an external mail service.

## Review the lab

- Message appears in the local inbox.
- Failed delivery is observable and does not deadlock requests.
- Durability limitations are documented.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
