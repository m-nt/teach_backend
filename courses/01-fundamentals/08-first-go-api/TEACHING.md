# Instructor notes: Building a small Go HTTP API

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. net/http calls a handler for a request. The handler translates HTTP input into Go values, invokes a rule, and translates the result into HTTP output. Keep transport details out of the core booking rule.

2. Decode JSON into a request struct, reject invalid values, and set the content type and status before writing the response. A missing error check can turn a bad request into a misleading success.

3. The example keeps data in process memory, so restarting loses bookings. HTTP handlers may execute concurrently; the supplied example locks its state. Students trace that lock today; concurrency mechanics belong to intermediate.

## Demonstration script

Trace POST /bookings in examples/03-http-api from JSON through validation and reserve to response. Change a field validation and send a failing request. Restart the process to demonstrate memory loss.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- What is a handler’s responsibility? **Expected reasoning:** Translate the request, call the rule, and write an HTTP response.
- Why validate before mutating state? **Expected reasoning:** A rejected request must not partially change the booking state.
- Will in-memory bookings survive a process restart? **Expected reasoning:** No.

## Review the lab

- /about returns JSON with HTTP 200.
- Invalid names produce 400 and do not consume seats.
- Restart behavior is explained correctly.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
