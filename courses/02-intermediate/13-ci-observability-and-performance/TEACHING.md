# Instructor notes: CI checks, structured logs, and measured performance

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. A continuous integration gate runs the same checks for every change: formatting, vet, tests, and a build. Include race checks and database integration tests where the runner supports them. Dependency vulnerability scanning complements tests.

2. Structured logs carry fields such as request ID, route, status, and duration. Do not log passwords, session tokens, OTPs, or unrestricted request bodies. Logs should answer what failed and where without exposing private data.

3. A benchmark measures a specific workload on a specific environment. Record dataset, concurrency, duration, error rate, and latency distribution before changing code. Use profiling to find a bottleneck rather than guessing.

## Demonstration script

Run the quality commands in resources/quality-gate.md. Trace a failed request across logs, then benchmark one rule and compare an intentional inefficient version.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- Does passing tests prove no known vulnerable dependencies exist? **Expected reasoning:** No; dependency scanning addresses a different concern.
- Why record error rate alongside latency? **Expected reasoning:** Fast failures can make latency appear better while service quality worsens.
- Why avoid full request bodies in logs? **Expected reasoning:** They can contain secrets and personal data.

## Review the lab

- CI fails for a broken test.
- A request ID links response evidence to logs.
- Measurement includes workload and environment details.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
