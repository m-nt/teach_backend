# Quiz: API contracts, pagination, and middleware

Allow 8 minutes, then 2 minutes for correction. Each question is worth 2 points: 1 for the key idea and 1 for a correct explanation or example. At least 4/6 is the lesson target; remediate missed prerequisites.

1. Is CORS an authorization mechanism?

2. Why order a paginated query?

3. Should public errors include database stack details?

<details>
<summary>Instructor answer key — open after submission</summary>

1. No; authorization must be enforced by the server.

2. To make page membership predictable.

3. No; log appropriate internal context and return a stable public error.

</details>
