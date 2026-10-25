# Instructor notes: Git, debugging, and a small project structure

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. Git records snapshots of tracked files. The working tree, staging area, and commit are different states. git status explains what is about to be included. A commit is local; a push sends it to a remote.

2. A Go module has a go.mod and one or more packages. Start with a single package while the program is small. Explain main.go, a rule file, and a test file; do not force a directory per concept.

3. Debug by reproducing, narrowing, changing one thing, and checking again. Read errors and use temporary prints or an editor breakpoint. Never commit credentials, environment files, or database dumps.

## Demonstration script

Initialize Git in a disposable copy of the CLI; inspect git diff, stage named files, commit, introduce an error, and repair it. Show how a second .go file in the same package is included by go run .

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- Is a commit the same as a push? **Expected reasoning:** No. A commit is local history; a push transfers commits to a remote.
- What does go.mod describe? **Expected reasoning:** The module path, Go directive, and module dependencies when present.
- Why inspect git diff before committing? **Expected reasoning:** To review exactly which changes are included and catch accidental content.

## Review the lab

- git log shows two meaningful commits.
- go run . works with multiple source files; no secrets are tracked.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
