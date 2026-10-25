# Quiz: Email adapters and bounded background work

Allow 8 minutes, then 2 minutes for correction. Each question is worth 2 points: 1 for the key idea and 1 for a correct explanation or example. At least 4/6 is the lesson target; remediate missed prerequisites.

1. Is an in-memory queue durable across restarts?

2. Does a timeout prove the email was not delivered?

3. Why use a fake sender in tests?

<details>
<summary>Instructor answer key — open after submission</summary>

1. No.

2. No, delivery may have happened before the timeout.

3. To verify behavior without depending on or contacting an external mail service.

</details>
