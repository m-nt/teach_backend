# Instructor notes: Goroutines, channels, mutexes, and bounded work

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. A goroutine runs concurrently with other goroutines. Shared memory needs coordination. A mutex protects an invariant across related reads and writes; locking only the write can still be incorrect.

2. Channels communicate values and completion. The sender normally owns closing a channel. Workers need a stop path; unbounded goroutines and abandoned sends can leak resources.

3. The race detector detects certain unsafe memory accesses during executed code. It does not prove transaction correctness or exercise every schedule. A double booking can be a business race even with no Go memory race.

## Demonstration script

Run an intentionally unsynchronized counter with go test -race, then protect the full update. Build a worker pool of two workers consuming five fake notifications.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- Does passing -race prove the absence of all races? **Expected reasoning:** No, it observes only exercised executions and addresses memory races.
- Who should normally close a channel? **Expected reasoning:** The sender or owner that knows no more sends will occur.
- Why bound concurrency? **Expected reasoning:** To control resource consumption and backpressure.

## Review the lab

- No race is reported in exercised paths.
- Cancellation test waits for every worker to stop.
- Worker count remains bounded.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
