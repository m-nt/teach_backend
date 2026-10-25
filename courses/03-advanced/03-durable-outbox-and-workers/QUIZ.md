# Quiz: Transactional outbox and durable notification workers

Allow 8 minutes, then 2 minutes for correction. Each question is worth 2 points: 1 for the key idea and 1 for a correct explanation or example. At least 4/6 is the lesson target; remediate missed prerequisites.

1. What failure gap does an outbox close?

2. Does an outbox automatically guarantee exactly-once email?

3. Why do claims need recovery?

<details>
<summary>Instructor answer key — open after submission</summary>

1. The gap between committing business state and separately recording delivery work.

2. No.

3. Workers may crash while holding work.

</details>
