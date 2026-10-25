# Lab: Password storage, sessions, and resource authorization

**Starting point:** your prior capstone commit and today’s prerequisite preparation.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Add users and sessions migrations and a maintained password-hash implementation.
2. Implement login, session lookup, expiry, and logout without logging credentials.
3. Enforce ownership on a booking endpoint and add a staff permission explicitly.
4. Test expired sessions, invalid passwords, cross-user access, and missing CSRF proof for browser mutations.

## Acceptance checks

- [ ] Cross-user access is denied even with a valid session.
- [ ] Logout invalidates the token.
- [ ] Stored password material is a password hash, not plaintext.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
