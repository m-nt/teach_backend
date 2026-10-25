# Instructor notes: Password storage, sessions, and resource authorization

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. Password hashing must use a maintained implementation of a password-specific algorithm with a unique salt and reviewed parameters. Follow the linked OWASP guidance; do not invent hashing or use a fast general hash alone.

2. For this course use opaque random server-side sessions. Store a digest of the token, enforce expiry and logout invalidation, and send browser cookies with HttpOnly, Secure in HTTPS environments, and an explicit SameSite policy. Add CSRF protection for cookie-authenticated state changes.

3. Authentication identifies a caller; ownership and role checks decide access. A student must not view or cancel another student’s booking. Test two accounts and a staff role; a guessed ID must not bypass authorization.

## Demonstration script

Trace the supplied design in resources/auth-design.md, implement one vertical login/session slice using the reviewed library, and demonstrate student A being denied student B’s booking.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- Why not store passwords using plain SHA-256? **Expected reasoning:** It is a fast general hash, not a password-specific work-factor scheme.
- Does login permit access to every resource? **Expected reasoning:** No, authorization is checked separately.
- Why does cookie authentication need CSRF consideration? **Expected reasoning:** Browsers can attach cookies automatically to requests initiated from another origin.

## Review the lab

- Cross-user access is denied even with a valid session.
- Logout invalidates the token.
- Stored password material is a password hash, not plaintext.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
