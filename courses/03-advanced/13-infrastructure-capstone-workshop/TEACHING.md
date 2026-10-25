# Instructor notes: Infrastructure capstone integration and rehearsal

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. The capstone integrates a business API, notification worker, PostgreSQL, Redis, sandbox mail, hostname routing, and verified TLS. Kubernetes runs the stateless workloads; document where stateful dependencies live and why.

2. Keep a deployment inventory with image versions, configuration names, ports, volumes, migrations, and recovery commands. A peer should be able to operate the system without your memory.

3. Use staged failures to test guarantees: replay a booking, race OTP verification, stop Redis, crash the worker, replace a Pod, and restore a backup. Fix core correctness before adding optional services.

## Demonstration script

Walk through CAPSTONE.md and assign one peer to operate each team’s runbook. Demonstrate one failure-and-recovery sequence without editing source during the demo.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- Why rehearse from a clean environment? **Expected reasoning:** To discover hidden local dependencies and undocumented steps.
- What proves a restore succeeded? **Expected reasoning:** Business-level data and application checks on the restored target.
- What should be prioritized before more services? **Expected reasoning:** Correctness, security, and recovery of the required system.

## Review the lab

- All required services and ownership boundaries are documented.
- Failures have observable, controlled outcomes.
- Remaining limitations are explicit.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
