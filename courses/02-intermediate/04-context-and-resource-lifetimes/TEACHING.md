# Instructor notes: Context, timeouts, defer, and resource lifetimes

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. A context carries cancellation, deadlines, and request-scoped values across calls. Pass it explicitly as the first parameter where appropriate. A request’s context should reach database and outbound HTTP operations.

2. A timeout limits waiting only when the called operation observes cancellation. Always invoke the cancel function returned by WithTimeout, usually using defer. Do not replace request context with Background deep inside the call chain.

3. defer runs when the surrounding function returns. Close response bodies and rows, check rows.Err, and avoid accumulating deferred cleanup inside a long-lived loop.

## Demonstration script

Build a fake slow HTTP dependency with httptest.Server. Give the caller a short deadline and observe cancellation. Compare a call using request context with one incorrectly using Background.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- Does a context forcibly kill any arbitrary function? **Expected reasoning:** No, the work must observe cancellation or use context-aware operations.
- Why call cancel even after success? **Expected reasoning:** To release context-associated resources promptly.
- When does defer execute? **Expected reasoning:** When the surrounding function returns.

## Review the lab

- Cancellation reaches the fake dependency.
- Timeout produces a controlled response and no extra mutation.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
