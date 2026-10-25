# Instructor notes: Timeouts, retries, backpressure, and overload

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. Timeouts consume an overall request budget. Retrying every layer multiplies calls; select one retry owner and retry only safe transient operations. Use bounded exponential backoff with jitter.

2. Backpressure limits admitted work when capacity is exhausted. Bound queue depth and worker concurrency, expose lag, and define a clear reject or defer response. More goroutines do not create unlimited downstream capacity.

3. A circuit breaker may reduce repeated calls to a failing dependency, but it adds state and recovery behavior. Compare it against a simpler timeout and concurrency limit before adopting it.

## Demonstration script

Use a fake dependency that fails three times and then succeeds. Plot attempts on a timeline; compare immediate unlimited retry with bounded backoff and a retry budget.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- Why can retries worsen an outage? **Expected reasoning:** They add load to an already failing or overloaded dependency.
- What does jitter help avoid? **Expected reasoning:** Many clients retrying at the same instant.
- Is every POST safe to retry? **Expected reasoning:** No; retry safety requires explicit idempotency or equivalent semantics.

## Review the lab

- Permanent failures are not retried indefinitely.
- Cancellation stops retry waits.
- Overload remains bounded and observable.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
