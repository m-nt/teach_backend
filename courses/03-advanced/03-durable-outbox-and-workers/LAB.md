# Lab: Transactional outbox and durable notification workers

**Starting point:** your prior capstone commit and today’s prerequisite preparation.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Add outbox and delivery receipt migrations.
2. Write booking and outbox in the same transaction.
3. Implement bounded claiming, lease recovery, and retry scheduling.
4. Crash before send and after send; record recovered jobs and duplicate behavior.

## Acceptance checks

- [ ] A committed booking always has delivery intent.
- [ ] Abandoned claims become processable again.
- [ ] Duplicate events do not duplicate internal business effects.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
