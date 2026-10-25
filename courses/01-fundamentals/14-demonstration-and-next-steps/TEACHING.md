# Instructor notes: Capstone demonstrations and fundamentals review

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. Assessment checks understanding and reproducibility, not the number of libraries used. Ask students to trace a request, explain a rule, and change one small behavior live.

2. Errors during a demo are useful evidence if the student diagnoses them methodically. Give a short repair opportunity and record the specific gap rather than assigning a vague label such as not technical.

3. The 20% milestone means introductory coverage across the local curriculum map. It does not certify production readiness. Intermediate will deepen Go, persistent booking correctness, security, and operational behavior.

## Demonstration script

Model a three-minute demonstration with one deliberate invalid request. Show how the rubric awards points for explaining limitations honestly.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- Where does the database fit in the request journey? **Expected reasoning:** The backend reads or writes it while handling business behavior.
- Why must a business rule be checked on the server? **Expected reasoning:** Clients can be changed or bypassed.
- What remains beyond fundamentals? **Expected reasoning:** Deeper language skills, concurrent persistent correctness, authentication, operations, and deployment.

## Review the lab

- Capstone rubric and evidence are recorded for each learner.
- Each learner has a concrete next practice task.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
