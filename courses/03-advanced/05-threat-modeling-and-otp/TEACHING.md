# Instructor notes: Threat modeling, OTP, and abuse controls

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. An OTP proves temporary access to a delivery channel; it is not automatically strong multi-factor authentication. Bind a challenge to purpose and account or destination. Generate codes using crypto/rand, not math/rand.

2. Short codes have low entropy. Store a keyed digest using a secret held outside the challenge store, expire quickly, limit attempts and issuance, and atomically verify-and-consume. Never log codes, tokens, or full destinations. Resending must not create unlimited concurrent valid codes.

3. Generic responses reduce account enumeration. Rate limits should combine relevant account/destination and network signals, with documented behavior when the limiter store fails. Do not use the public-catalog cache fallback policy for authentication controls.

## Demonstration script

Use resources/auth-design.md to trace issue → sandbox delivery → verify → consume. Attempt the same code twice and race two verification requests.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- Why is plain hashing of a short numeric code insufficient against offline guessing? **Expected reasoning:** The small code space is cheap to enumerate; a keyed digest adds protection when the store alone is exposed.
- Why consume atomically? **Expected reasoning:** Concurrent verifications must not both accept the same challenge.
- Is email OTP automatically phishing-resistant MFA? **Expected reasoning:** No.

## Review the lab

- Only one concurrent verification succeeds.
- Wrong attempts and issuance are bounded.
- No code or session token appears in logs.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
