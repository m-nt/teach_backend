# Instructor notes: Linux operations, network boundaries, and secrets

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. A process runs as a user and listens on selected interfaces. Bind public entrypoints deliberately; databases and Redis should be reachable only by intended applications. Explain loopback, private network, and public interface with a packet path.

2. Use SSH keys, least-privilege accounts, and a documented firewall policy. Keep an existing administrative session open while testing access changes to avoid locking yourself out.

3. Secrets need controlled injection, rotation, and restricted access. Environment variables and Kubernetes Secrets are delivery mechanisms, not automatic protection against every reader or log leak. Keep secrets out of images, source control, and command transcripts.

## Demonstration script

On a disposable VM inspect processes and listening sockets with ps and ss. Trace a request to the reverse proxy and show that the database is not publicly exposed. Demonstrate configuration using placeholder secrets.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- Does running in a container automatically isolate every published port? **Expected reasoning:** No; published ports and host firewall rules still matter.
- Why keep an existing SSH session during access changes? **Expected reasoning:** To preserve a recovery path if new access fails.
- Is a base64-encoded secret encrypted? **Expected reasoning:** No.

## Review the lab

- Only intended entrypoints are exposed.
- Rotation works without a source-code edit.
- Recovery access is tested before closing the original session.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
