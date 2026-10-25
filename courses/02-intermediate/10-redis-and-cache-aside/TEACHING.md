# Instructor notes: Redis and cache-aside behavior

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. A cache holds a derived copy of data. In cache-aside, read the cache, load the database on a miss, and populate with a bounded TTL. The database remains the source of truth.

2. Invalidation matters more than the happy-path hit. Delete or refresh affected keys after a successful write. A stale capacity display must never authorize a booking; the transactional rule still checks the database.

3. Cache failure should not silently change correctness. For public catalog reads, fall back to the database. Scope keys correctly and do not cache private responses under a shared public key.

## Demonstration script

Use redis-cli SET and GET with expiry, then trace cold and warm GET /sessions requests. Stop Redis and demonstrate a successful database fallback.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- What is the source of truth in this design? **Expected reasoning:** PostgreSQL.
- Does a TTL eliminate all stale reads? **Expected reasoning:** No; data can be stale until expiry or invalidation.
- Why are shared private-response cache keys dangerous? **Expected reasoning:** They can return one user’s data to another user.

## Review the lab

- An update becomes visible after invalidation.
- Booking authority never depends on the cache.
- Redis outage does not break the public read.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
