# Lab: Metrics, traces, SLOs, and incident diagnosis

**Starting point:** your prior capstone commit and today’s prerequisite preparation.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Define one availability and one latency SLO with a measurement window.
2. Expose bounded metrics and propagate trace context across the notification message.
3. Trigger a controlled dependency incident and collect evidence.
4. Write an incident timeline, mitigation, and follow-up action.

## Acceptance checks

- [ ] Telemetry distinguishes API success from delayed email delivery.
- [ ] No secret or user ID is used as an unbounded metric label.
- [ ] Incident diagnosis cites observations.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
