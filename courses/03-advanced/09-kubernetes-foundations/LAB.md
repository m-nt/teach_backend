# Lab: Kubernetes: Pods, Deployments, Services, and configuration

**Starting point:** your prior capstone commit and today’s prerequisite preparation.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Create a course namespace and configuration using placeholders replaced locally.
2. Load or publish the student image and apply API/worker Deployments.
3. Expose the API through a ClusterIP Service and kubectl port-forward.
4. Delete one Pod, inspect events, and watch the controller restore it.

## Acceptance checks

- [ ] API and worker run as separate Deployments.
- [ ] Service reaches replacement Pods.
- [ ] No real secret is committed to the manifests.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
