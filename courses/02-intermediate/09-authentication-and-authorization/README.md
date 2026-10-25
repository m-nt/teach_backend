# Intermediate 09: Password storage, sessions, and resource authorization

[Tier index](../README.md) · [Instructor notes](TEACHING.md) · [Lab](LAB.md) · [Quiz](QUIZ.md) · [Homework](HOMEWORK.md)

**Duration:** 180 minutes. **Outcome:** Implement reviewed password hashing and session-based access with ownership checks.

## Prerequisites

Complete [API contracts, pagination, and middleware](../08-api-contracts-and-middleware/README.md), including its next-class preparation.

## Class plan

| Minutes | Activity |
|---|---|
| 00–15 | Retrieval practice and prerequisite check |
| 15–45 | Concepts, design discussion, and whiteboard |
| 45–70 | Live demonstration and failure case |
| 70–80 | Break |
| 80–135 | Guided implementation lab |
| 135–160 | Failure tests and peer review |
| 160–170 | Quiz and correction |
| 170–180 | Exit check and next-class preparation |

## Lesson

### 1. Password hashing must use a maintained implementation of a password-specific algorithm with a unique salt and reviewed parameters.

Follow the linked OWASP guidance; do not invent hashing or use a fast general hash alone.

### 2. For this course use opaque random server-side sessions.

Store a digest of the token, enforce expiry and logout invalidation, and send browser cookies with HttpOnly, Secure in HTTPS environments, and an explicit SameSite policy. Add CSRF protection for cookie-authenticated state changes.

### 3. Authentication identifies a caller; ownership and role checks decide access.

A student must not view or cancel another student’s booking. Test two accounts and a staff role; a guessed ID must not bypass authorization.

## Exit check

- Cross-user access is denied even with a valid session.
- Logout invalidates the token.
- Stored password material is a password hash, not plaintext.

## References

Use the topic-matched official references in [resources/sources.md](../../../resources/sources.md). The explanation and exercise here are original course material; the roadmaps are sequencing references, not assigned reading in full.
