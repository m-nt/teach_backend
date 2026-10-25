# Lab: Building a small Go HTTP API

**Starting point:** the referenced repository example or an empty personal exercise folder.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Copy the HTTP example to a personal exercise folder.
2. Add GET /about with a JSON course description.
3. Add a rule rejecting student names shorter than two characters after trimming.
4. Exercise valid, invalid, absent-session, and full-session requests using the contract.

## Acceptance checks

- [ ] /about returns JSON with HTTP 200.
- [ ] Invalid names produce 400 and do not consume seats.
- [ ] Restart behavior is explained correctly.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
