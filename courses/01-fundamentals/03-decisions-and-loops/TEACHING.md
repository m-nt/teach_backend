# Instructor notes: Decisions, loops, and tracing

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. A condition evaluates to true or false. A booking is allowed only when the class is open and remaining seats are positive. Translate the sentence into small conditions before combining them with &&.

2. A loop repeats work. Trace one iteration at a time on paper with columns for attempt and remaining seats. The loop condition is checked before each iteration; updating the wrong variable can prevent termination.

3. Boundary values expose mistakes: zero seats, exactly one seat, and a closed class with seats available. switch is useful for named states; it does not replace all if statements.

## Demonstration script

Use a two-seat class and simulate three booking requests in a for loop. Predict each output before running. Show an off-by-one error with <= and fix it.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- What does && require? **Expected reasoning:** Both conditions must be true.
- Why test zero and one seat? **Expected reasoning:** They exercise the boundary where the booking decision changes.
- Why must the loop state change? **Expected reasoning:** So a bounded loop can progress toward its stopping condition.

## Review the lab

- Five attempts accept exactly three and reject two.
- A closed class accepts no requests.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
