# Instructor notes: Delivery pipelines, release rollback, and safe migrations

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. Build once and promote the same immutable artifact. Record source revision and image digest; floating tags make incident reconstruction harder. Run quality checks before promotion.

2. Use expand/migrate/contract for schema changes that overlap old and new application versions. A rollback of code does not automatically reverse data changes. Backups do not replace a compatibility plan.

3. Deployment evidence includes health, a synthetic transaction, error metrics, and an explicit rollback trigger. Migration jobs should run once under controlled ownership, not independently in every replica.

## Demonstration script

Deploy version A, apply an additive schema migration, deploy B, then roll back to A against the changed schema. Show why dropping a required old column would break that rollback.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- Does rolling back an image undo a schema migration? **Expected reasoning:** No.
- Why deploy the same image artifact across environments? **Expected reasoning:** To reduce differences between what was tested and what runs.
- When should a destructive schema contraction happen? **Expected reasoning:** After old consumers are retired and compatibility/recovery requirements are met.

## Review the lab

- Release evidence names a digest and source revision.
- Rollback completes without losing bookings.
- Migration ownership is explicit.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
