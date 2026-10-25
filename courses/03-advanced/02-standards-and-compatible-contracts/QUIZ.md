# Quiz: Standards, contracts, and compatible evolution

Allow 8 minutes, then 2 minutes for correction. Each question is worth 2 points: 1 for the key idea and 1 for a correct explanation or example. At least 4/6 is the lesson target; remediate missed prerequisites.

1. Is a new optional field always safe for every consumer?

2. Why scope idempotency keys to a caller?

3. What must a replay return?

<details>
<summary>Instructor answer key — open after submission</summary>

1. No; safety depends on the documented parsing behavior and tests.

2. To prevent collisions or cross-user result exposure.

3. The documented outcome of the original logical operation, not a second side effect.

</details>
