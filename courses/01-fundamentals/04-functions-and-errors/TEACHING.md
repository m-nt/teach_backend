# Instructor notes: Functions, parameters, and explicit errors

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. A function names a task. Parameters are inputs and return values are outputs. Start with remaining(capacity, booked int) int, then extract a rule that can fail.

2. Go commonly returns a value and an error. nil means no error. A caller should check the error before using a result that may be invalid. Returning an error lets the caller choose whether to print a message or later send an HTTP response.

3. Keep business rules separate from presentation. A booking function should return a result rather than printing every decision. That same function can serve a CLI now and an HTTP handler later.

## Demonstration script

Type remaining and reserve functions. Have reserve reject booked >= capacity with fmt.Errorf. Call it for a normal class and a full class, showing both branches.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- What does nil error mean? **Expected reasoning:** The operation did not report an error.
- Why return an error instead of always printing? **Expected reasoning:** The caller can decide how to present or handle the failure.
- Are parameters declared in one function visible in another? **Expected reasoning:** No, each function has its own local scope.

## Review the lab

- A full class returns an error and keeps the original state.
- A valid request increments booked once.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
