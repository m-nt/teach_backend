# Lab: Threat modeling, OTP, and abuse controls

**Starting point:** your prior capstone commit and today’s prerequisite preparation.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Implement a local verification challenge with purpose binding and an expiry.
2. Add keyed digest storage, attempt limits, resend cooldown, and atomic consumption.
3. Send only to Mailpit and test wrong, expired, reused, and concurrent codes.
4. Stop Redis or the selected challenge store and confirm the documented authentication failure policy.

## Acceptance checks

- [ ] Only one concurrent verification succeeds.
- [ ] Wrong attempts and issuance are bounded.
- [ ] No code or session token appears in logs.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
