# Quiz: Kubernetes probes, resources, rollout, and network controls

Allow 8 minutes, then 2 minutes for correction. Each question is worth 2 points: 1 for the key idea and 1 for a correct explanation or example. At least 4/6 is the lesson target; remediate missed prerequisites.

1. Should liveness fail whenever PostgreSQL is briefly unavailable?

2. Does declaring a NetworkPolicy ensure enforcement on every cluster?

3. Can CPU autoscaling fix all latency problems?

<details>
<summary>Instructor answer key — open after submission</summary>

1. Usually no; readiness and controlled request failures can express that dependency state.

2. No, the network implementation must support it.

3. No; bottlenecks may be elsewhere.

</details>
