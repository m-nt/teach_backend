# Lab: Email adapters and bounded background work

**Starting point:** your prior capstone commit and today’s prerequisite preparation.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Define an email adapter and a fake sender.
2. Render a simple confirmation without sensitive authentication data.
3. Use a bounded worker and context-aware shutdown.
4. Test failure and record which jobs are lost when the process stops.

## Acceptance checks

- [ ] Message appears in the local inbox.
- [ ] Failed delivery is observable and does not deadlock requests.
- [ ] Durability limitations are documented.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
