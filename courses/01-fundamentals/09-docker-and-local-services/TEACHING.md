# Instructor notes: Docker, images, containers, ports, and volumes

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. An image packages a filesystem and launch configuration; a container is a running instance. Docker is a tool for running dependencies consistently, not a replacement for understanding them.

2. A published port connects a host port to a container port. localhost inside a container refers to that container. Compose service names identify other services on its network. Keep the training database bound to the host loopback interface.

3. A volume stores data separately from a container lifecycle. Stopping or replacing a container is different from deleting its volume. The local Compose file uses development credentials and must stay local.

## Demonstration script

From infrastructure/local run docker compose up -d db, then docker compose ps and logs. Enter psql with the supplied command. Stop/start the container and show the volume remains.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- Image versus container? **Expected reasoning:** A packaged template versus an instance running from it.
- What does localhost mean inside a container? **Expected reasoning:** That container itself.
- Does stopping a container necessarily delete its volume? **Expected reasoning:** No.

## Review the lab

- SELECT 1 returns one row.
- Student identifies host port 5432 and container port 5432.
- Student explains where database files persist.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
