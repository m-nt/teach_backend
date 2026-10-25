# Instructor notes: Standards, contracts, and compatible evolution

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. Use consistent HTTP semantics and a documented error shape. OpenAPI describes synchronous APIs; a versioned event schema describes asynchronous messages. Standards reduce ambiguity but do not choose business rules for you.

2. An event needs an ID, type, schema version, occurrence time, and business reference. Consumers should tolerate documented additive fields. Removing or changing a field requires a migration strategy.

3. Idempotency means a repeated logical operation has the intended single effect. Define a client key’s scope, request fingerprint, stored response, and expiry. Reusing a key for a different request should be rejected.

## Demonstration script

Compare version 1 and an additive version 2 notification payload. Replay a booking request with an idempotency key, then reuse the key with a different payload and show a conflict.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- Is a new optional field always safe for every consumer? **Expected reasoning:** No; safety depends on the documented parsing behavior and tests.
- Why scope idempotency keys to a caller? **Expected reasoning:** To prevent collisions or cross-user result exposure.
- What must a replay return? **Expected reasoning:** The documented outcome of the original logical operation, not a second side effect.

## Review the lab

- Repeated requests create one booking and return a consistent result.
- Different payload with the same key is rejected.
- Compatibility expectations are executable or explicitly demonstrated.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
