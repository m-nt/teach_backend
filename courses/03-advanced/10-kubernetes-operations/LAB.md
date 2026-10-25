# Lab: Kubernetes probes, resources, rollout, and network controls

**Starting point:** your prior capstone commit and today’s prerequisite preparation.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Set probes, requests/limits, and graceful termination for the API.
2. Perform a rollout and observe readiness before traffic.
3. Apply a limited service account and a tested network policy on a supporting cluster.
4. Diagnose one intentionally broken image or configuration and recover.

## Acceptance checks

- [ ] Unready Pods are excluded from Service endpoints.
- [ ] The failure is explained using events or logs.
- [ ] Network denial is verified or explicitly marked unverified if the plugin lacks support.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
