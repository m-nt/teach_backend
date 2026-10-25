# Instructor notes: Go value semantics, pointers, and methods

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. Assigning a struct copies its fields. A pointer refers to an existing value, so changing through the pointer can affect the caller’s value. Explain ownership using two boxes and an arrow before showing syntax.

2. A slice value describes a backing array; copying the slice does not necessarily copy its elements. append may allocate another array. Code should not depend on accidental spare capacity.

3. Methods associate behavior with a type. Choose pointer receivers when mutation or consistent receiver behavior requires them. A pointer is not automatically faster or safer.

## Demonstration script

Create a Session value and pass it to a function by value, then by pointer. Repeat with a slice and show shared element mutation. Draw the objects before running.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- Does assigning a struct create an alias to the original struct? **Expected reasoning:** No, it copies the struct value.
- Does copying a slice guarantee independent elements? **Expected reasoning:** No, it can share the backing array.
- When is a pointer receiver appropriate? **Expected reasoning:** When the method must mutate its receiver or receiver consistency requires it.

## Review the lab

- Prediction matches observed struct and slice behavior.
- Tests prove Reserve mutates only the intended value.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
