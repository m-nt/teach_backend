# Instructor notes: API contracts, pagination, and middleware

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. A contract specifies paths, methods, schemas, statuses, and failure shapes. OpenAPI can document it, but generated text still needs review. Avoid exposing internal errors as public messages.

2. Pagination needs deterministic ordering. Start with bounded limit/offset for the small course catalog; explain how concurrent changes and large offsets motivate cursor pagination later.

3. Middleware handles transport-wide work such as request IDs and logs. Validate input at the boundary and authorize each protected operation. CORS config controls browser access; it is not an authorization system.

## Demonstration script

Write an OpenAPI fragment for GET /sessions and POST /bookings. Add a maximum page size and request ID. Send an oversized limit and inspect the documented response.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- Is CORS an authorization mechanism? **Expected reasoning:** No; authorization must be enforced by the server.
- Why order a paginated query? **Expected reasoning:** To make page membership predictable.
- Should public errors include database stack details? **Expected reasoning:** No; log appropriate internal context and return a stable public error.

## Review the lab

- Documented examples match actual responses.
- Page order is deterministic and size is bounded.
- Logs include request IDs without request secrets.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
