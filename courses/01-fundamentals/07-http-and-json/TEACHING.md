# Instructor notes: HTTP, URLs, JSON, and API contracts

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. HTTP connects a client request with a server response. A URL contains a scheme, host, optional port, path, and possibly a query string. A method expresses intent; headers carry metadata; the body carries data.

2. JSON represents objects, arrays, strings, numbers, booleans, and null. It is a wire format, not a database. Distinguish a valid JSON document from a document with valid business values.

3. Agree on a contract before implementing handlers: GET /sessions lists classes; POST /bookings requests a booking. Use 200 for a successful read, 201 for creation, 400 for malformed input, 404 for absent resources, and 409 for a full class conflict. These are this course’s contract choices.

## Demonstration script

Run examples/03-http-api and inspect curl -i http://localhost:8080/sessions. Read the status line, content type, and JSON separately. Send malformed JSON to POST /bookings.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- Is HTTP 200 a JSON field? **Expected reasoning:** No. It is an HTTP response status.
- Can syntactically valid JSON still be rejected? **Expected reasoning:** Yes, required fields or business rules can be invalid.
- Why should a full class have a predictable error response? **Expected reasoning:** Clients need a stable contract to display and handle the failure.

## Review the lab

- Student identifies method, path, status, headers, and body.
- Contract contains success and failure examples.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
