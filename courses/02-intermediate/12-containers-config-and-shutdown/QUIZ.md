# Quiz: Containerizing Go, configuration, and graceful shutdown

Allow 8 minutes, then 2 minutes for correction. Each question is worth 2 points: 1 for the key idea and 1 for a correct explanation or example. At least 4/6 is the lesson target; remediate missed prerequisites.

1. Why is localhost wrong for another Compose service?

2. Should temporary database failure always fail liveness?

3. Why bound graceful shutdown?

<details>
<summary>Instructor answer key — open after submission</summary>

1. It points to the current container, not the dependency.

2. No; restarting every process can amplify the outage. Readiness can reflect dependency availability.

3. To avoid waiting forever for stuck work.

</details>
