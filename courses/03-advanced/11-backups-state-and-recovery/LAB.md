# Lab: Stateful services, backups, and restore drills

**Starting point:** your prior capstone commit and today’s prerequisite preparation.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Record RPO/RTO targets and storage ownership.
2. Create a consistent database backup and protect its access.
3. Restore to an isolated target without overwriting the source.
4. Verify row counts, relationships, known bookings, and API behavior; measure elapsed recovery time.

## Acceptance checks

- [ ] Restored data is validated through SQL and the API.
- [ ] Source database remains untouched.
- [ ] Observed data loss and recovery duration are compared to targets.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
