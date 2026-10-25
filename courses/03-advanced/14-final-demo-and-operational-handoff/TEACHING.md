# Instructor notes: Final demonstration and operational handoff

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. The final demonstration shows working behavior and recovery, not only diagrams. Require evidence for TLS verification, authorization, concurrency correctness, durable work, and data restoration.

2. An operational handoff names owners, runbooks, release procedure, secrets rotation, backups, alerts, and known risks. A system nobody can recover is not finished merely because it starts.

3. Completion means students can build and operate the defined small infrastructure with review. It does not imply exhaustive backend expertise or readiness to operate every workload without supervision.

## Demonstration script

Model a short handoff: architecture, verified user flow, failure, recovery, and remaining risk. Use the rubric consistently across teams.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- Does a green dashboard prove every business invariant? **Expected reasoning:** No; targeted correctness tests and checks remain necessary.
- Why document operational ownership? **Expected reasoning:** Incidents and maintenance need accountable action.
- What follows this course? **Expected reasoning:** Practice against real requirements, supervised operation, and deeper specialization based on measured needs.

## Review the lab

- Mandatory advanced safety and correctness gates pass.
- A peer can follow the operational runbook.
- Each student explains a subsystem they personally implemented.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
