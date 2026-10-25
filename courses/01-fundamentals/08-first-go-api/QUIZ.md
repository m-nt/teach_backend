# Quiz: Building a small Go HTTP API

Allow 8 minutes, then 2 minutes for correction. Each question is worth 2 points: 1 for the key idea and 1 for a correct explanation or example. At least 4/6 is the lesson target; remediate missed prerequisites.

1. What is a handler’s responsibility?

2. Why validate before mutating state?

3. Will in-memory bookings survive a process restart?

<details>
<summary>Instructor answer key — open after submission</summary>

1. Translate the request, call the rule, and write an HTTP response.

2. A rejected request must not partially change the booking state.

3. No.

</details>
