# Instructor notes: Metrics, traces, SLOs, and incident diagnosis

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. An SLI is a measured indicator; an SLO is its target over a stated window. Define successful booking requests and acceptable latency precisely, including exclusions. A single successful probe is not a service-level history.

2. Metrics reveal rates, errors, durations, saturation, and queue lag. Keep labels bounded; user IDs and arbitrary URLs create excessive cardinality. Logs explain events; traces connect work across API, storage, and worker boundaries.

3. Alert on actionable user impact and clear ownership. Use a runbook to identify whether the incident is in DNS, proxy, API, database, cache, queue, or mail delivery. Measure before changing multiple things.

## Demonstration script

Inject database latency or pause the notification worker. Compare API latency, errors, and queue age. Follow one request ID or trace and identify the actual failing boundary.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- SLI versus SLO? **Expected reasoning:** The measurement versus its target over a defined window.
- Why avoid user IDs as metric labels? **Expected reasoning:** They create unbounded cardinality and can expose private information.
- Can HTTP success alone prove notification delivery? **Expected reasoning:** No, asynchronous delivery needs its own indicator.

## Review the lab

- Telemetry distinguishes API success from delayed email delivery.
- No secret or user ID is used as an unbounded metric label.
- Incident diagnosis cites observations.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
