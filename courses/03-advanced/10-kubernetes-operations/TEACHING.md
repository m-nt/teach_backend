# Instructor notes: Kubernetes probes, resources, rollout, and network controls

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. Readiness controls whether a Pod receives normal Service traffic; liveness can restart a container. A startup probe allows slow initialization. Dependency outages should not cause every healthy process to restart continuously.

2. Set requests and limits based on observation. Inspect events for scheduling failures and terminated-container reasons for memory exhaustion. Autoscaling needs appropriate metrics and does not fix a database bottleneck.

3. NetworkPolicies require a supporting network plugin. Restrict ingress/egress by the actual dependency graph, then test both allowed and denied flows. RBAC limits API actions; it is different from network access.

## Demonstration script

Break readiness, request excessive resources, and inspect describe/events/logs. Restore the configuration and watch a rolling update. Explain how a non-enforcing network plugin would invalidate a policy demo.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- Should liveness fail whenever PostgreSQL is briefly unavailable? **Expected reasoning:** Usually no; readiness and controlled request failures can express that dependency state.
- Does declaring a NetworkPolicy ensure enforcement on every cluster? **Expected reasoning:** No, the network implementation must support it.
- Can CPU autoscaling fix all latency problems? **Expected reasoning:** No; bottlenecks may be elsewhere.

## Review the lab

- Unready Pods are excluded from Service endpoints.
- The failure is explained using events or logs.
- Network denial is verified or explicitly marked unverified if the plugin lacks support.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
