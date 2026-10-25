# Instructor notes: Fundamentals capstone workshop

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. A small complete slice is more valuable than many disconnected endpoints. The core capstone combines a session catalog with persistent student registration; booking rules can remain in a tested CLI. Persistent, concurrent booking is an intermediate outcome.

2. A run guide is part of the deliverable. It states prerequisites, startup order, sample requests, and known limitations. Another person should be able to reproduce results without your terminal history.

3. Review one behavior at a time: valid input, invalid input, persistence, and dependency failure. Keep work visible using a short checklist; use the workshop to repair gaps rather than add new infrastructure.

## Demonstration script

Demonstrate the capstone acceptance flow: start db, run API, list sessions/students, create a student, reject an invalid name, restart, and run rule tests. Point out the boundary between supplied scaffolding and student changes.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- What is an acceptance criterion? **Expected reasoning:** An observable condition showing a requirement is satisfied.
- Does a working happy path prove a backend is correct? **Expected reasoning:** No; failure and boundary paths also need checking.
- What should a run guide include? **Expected reasoning:** Prerequisites, startup, verification requests, tests, and limitations.

## Review the lab

- Partner can start the project using its guide.
- Persistence and invalid-input behavior are demonstrated.
- Student identifies which code they wrote and which scaffold they used.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
