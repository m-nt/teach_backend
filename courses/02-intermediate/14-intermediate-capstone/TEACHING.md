# Instructor notes: Intermediate capstone: a reliable single-service backend

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. This milestone deepens the same project rather than rewarding the number of endpoints. The hardest checks are simultaneous reservation, ownership enforcement, and behavior when dependencies fail.

2. A demonstration is reproducible evidence. Keep scripts and expected outcomes together, reset fixtures safely, and explain why a test targets a real failure mode.

3. The 50% milestone means completing the first two tiers’ coverage goals. Advanced addresses distributed failure, durable notifications, infrastructure, domain routing, TLS, and operational recovery.

## Demonstration script

Model the last-seat concurrency test and a denied cross-user cancellation. Demonstrate a cache outage with a working database fallback.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- Which system decides booking capacity? **Expected reasoning:** The transactional database-backed booking logic.
- What does a passing local stack not prove? **Expected reasoning:** Public deployment security, disaster recovery, and distributed reliability.
- Why retain explicit limitations? **Expected reasoning:** They define the real guarantees and guide the next improvements.

## Review the lab

- All mandatory intermediate gates pass.
- Evidence includes tests and live failure demonstrations.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
