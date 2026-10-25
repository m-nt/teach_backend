# Instructor notes: Containerizing Go, configuration, and graceful shutdown

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. A multi-stage build compiles Go in a build stage and copies the binary into a smaller runtime stage. Choose a runtime that meets the program’s needs, such as CA certificates for outbound TLS.

2. Configuration should be validated at startup. Missing required secrets or malformed addresses should fail clearly. Containers on a Compose network address dependencies by service name, not host localhost.

3. Graceful shutdown stops accepting new work, lets bounded in-flight work finish, and closes resources. Readiness means able to serve; liveness means the process should remain running. Their checks need different failure behavior.

## Demonstration script

Build the student API image using resources/Dockerfile.example. Run it with db and Redis on a Compose network. Send SIGTERM during a controlled slow request and observe bounded shutdown.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- Why is localhost wrong for another Compose service? **Expected reasoning:** It points to the current container, not the dependency.
- Should temporary database failure always fail liveness? **Expected reasoning:** No; restarting every process can amplify the outage. Readiness can reflect dependency availability.
- Why bound graceful shutdown? **Expected reasoning:** To avoid waiting forever for stuck work.

## Review the lab

- Image starts without local source files.
- API connects using Compose service names.
- Shutdown finishes or times out predictably.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
