# Quiz: Redis and cache-aside behavior

Allow 8 minutes, then 2 minutes for correction. Each question is worth 2 points: 1 for the key idea and 1 for a correct explanation or example. At least 4/6 is the lesson target; remediate missed prerequisites.

1. What is the source of truth in this design?

2. Does a TTL eliminate all stale reads?

3. Why are shared private-response cache keys dangerous?

<details>
<summary>Instructor answer key — open after submission</summary>

1. PostgreSQL.

2. No; data can be stale until expiry or invalidation.

3. They can return one user’s data to another user.

</details>
