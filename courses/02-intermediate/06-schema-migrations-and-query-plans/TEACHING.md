# Instructor notes: Schema migrations, indexes, and query plans

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. A migration records an ordered schema change. Separate schema changes from ad hoc edits, and test both an empty database and an upgrade with existing data. Rolling deployments may run old and new application versions together.

2. An index trades storage and write work for faster access patterns. Choose indexes from WHERE, JOIN, and ORDER BY requirements. A tiny table may reasonably use a sequential scan.

3. Use EXPLAIN and EXPLAIN ANALYZE to inspect plans; ANALYZE executes the statement. Use synthetic read queries for this lab and compare results at a meaningful data size.

## Demonstration script

Create an additive migration for session start time. Populate synthetic sessions, compare a filtered ordered query before and after a matching index, and discuss why a small fixture can mislead.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- Does every column need an index? **Expected reasoning:** No; indexes should serve measured access patterns.
- Does EXPLAIN ANALYZE execute its query? **Expected reasoning:** Yes.
- Why test migrations with old data? **Expected reasoning:** Constraints and data transformations can fail despite a clean install working.

## Review the lab

- Existing booking data survives the upgrade.
- Plan comparison includes row count and query text.
- Rollback data-loss consequences are stated.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
