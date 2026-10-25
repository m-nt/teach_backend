# Authentication and OTP implementation design

This is a teaching specification and failure-test guide, not a ready-made authentication library. Use the maintained implementations and current guidance in [sources.md](sources.md). Rehearse the selected library integration before teaching intermediate lesson 09.

## Password login and sessions

Use Argon2id through a maintained implementation, with parameters reviewed against current OWASP guidance and measured on the lab hardware. Store algorithm/version/parameters, unique salt, and hash in an encoded record. Compare using the library’s verification function; reject malformed records safely. Never store plaintext or use plain SHA-256 for passwords.

A minimal schema adds `users(id, email UNIQUE, password_hash, role)` and `sessions(token_digest UNIQUE, user_id REFERENCES users, expires_at, created_at)`. Add a user-to-student relationship with a unique constraint if they are separate entities. Generate at least 32 random token bytes using `crypto/rand`; encode for transport, and store only a cryptographic digest of the high-entropy session token. This differs from a short numeric OTP, where a plain digest can be exhaustively guessed.

Login sequence: bounded input → normalize only appropriate fields → rate-limit attempt → lookup → verify → create expiring session → send cookie. Use generic invalid-credential responses. For a missing user, use the chosen library’s dummy verification pattern to reduce obvious timing differences. Restrict roles to server-owned data.

For browsers choose HttpOnly, Secure under HTTPS, an explicit SameSite policy, and a host-only cookie scope. Use a server-generated CSRF token checked on mutations, with origin checks as additional protection. For local plain-HTTP exercises, any Secure-cookie exception must be limited to the local development configuration. Do not silently ship that exception. Logout removes the server-side session. Each protected request validates expiry and ownership independently of login.

Test matrix: correct/wrong password, absent account, expired token, logout then reuse, user A reading/cancelling user B’s booking, missing/invalid CSRF token, oversized input, rate-limit threshold, missing required secret. Secrets and authentication bodies must be absent from logs.

## OTP verification challenge

For the lab choose a purpose such as verifying a synthetic email address. The following are explicit **classroom design choices**, not universal recommended values: six digits, five-minute lifetime, five wrong attempts maximum, 60-second resend cooldown, and a bounded issuance budget per account/destination plus a separate network budget. Review the threat model before using any of them outside the lab.

State: `challenge_id`, `subject`, `purpose`, `code_digest`, `expires_at`, `attempts_remaining`, and `consumed_at`. Use a random public challenge ID. Generate the code with unbiased `crypto/rand` sampling. Compute a keyed HMAC over the challenge ID, purpose, subject, and code with unambiguous encoding; keep the HMAC key outside the challenge store. Do not log the code or include it in the response. Deliver only through Mailpit in the course.

Verification transaction / atomic operation:

1. Load and lock the challenge or use a carefully reviewed atomic Redis script.
2. Reject absent, expired, consumed, wrong-purpose, or exhausted challenges.
3. On mismatch, decrement attempts and commit that decrement. **Do not roll back the failed-attempt count.**
4. On a valid match, atomically mark consumed and record the allowed verification result.
5. Only one concurrent caller may observe a usable challenge and consume it successfully.

Issuance and resend must serialize their shared budget/state; invalidating the prior code and recording the new one must not race. Use generic public responses. If the challenge/rate-limit store fails, deny or defer verification with a controlled retryable error; do not fall through as authenticated. A public-read cache fallback is not suitable here.

Required tests: wrong, expired, replayed, exhausted, wrong-purpose, simultaneous correct verification, concurrent resends, issue budget, store outage, and secret-free logs. Explain that email OTP is not phishing-resistant MFA and that an account recovery design needs separate review.
