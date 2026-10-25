# Quiz: Timeouts, retries, backpressure, and overload

Allow 8 minutes, then 2 minutes for correction. Each question is worth 2 points: 1 for the key idea and 1 for a correct explanation or example. At least 4/6 is the lesson target; remediate missed prerequisites.

1. Why can retries worsen an outage?

2. What does jitter help avoid?

3. Is every POST safe to retry?

<details>
<summary>Instructor answer key — open after submission</summary>

1. They add load to an already failing or overloaded dependency.

2. Many clients retrying at the same instant.

3. No; retry safety requires explicit idempotency or equivalent semantics.

</details>
