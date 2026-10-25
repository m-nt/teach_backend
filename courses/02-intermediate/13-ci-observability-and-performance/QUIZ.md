# Quiz: CI checks, structured logs, and measured performance

Allow 8 minutes, then 2 minutes for correction. Each question is worth 2 points: 1 for the key idea and 1 for a correct explanation or example. At least 4/6 is the lesson target; remediate missed prerequisites.

1. Does passing tests prove no known vulnerable dependencies exist?

2. Why record error rate alongside latency?

3. Why avoid full request bodies in logs?

<details>
<summary>Instructor answer key — open after submission</summary>

1. No; dependency scanning addresses a different concern.

2. Fast failures can make latency appear better while service quality worsens.

3. They can contain secrets and personal data.

</details>
