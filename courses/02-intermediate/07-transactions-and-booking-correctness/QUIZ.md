# Quiz: Transactions and concurrent booking correctness

Allow 8 minutes, then 2 minutes for correction. Each question is worth 2 points: 1 for the key idea and 1 for a correct explanation or example. At least 4/6 is the lesson target; remediate missed prerequisites.

1. Why is a pre-check outside the transaction insufficient?

2. What does the unique pair constraint prevent?

3. Why keep network calls outside the transaction?

<details>
<summary>Instructor answer key — open after submission</summary>

1. Another caller can change availability between the check and write.

2. Two bookings for the same student and session.

3. Slow or failed calls prolong locks and couple unrelated failures.

</details>
