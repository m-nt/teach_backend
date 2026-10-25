# Quiz: Table-driven tests, HTTP tests, and integration boundaries

Allow 8 minutes, then 2 minutes for correction. Each question is worth 2 points: 1 for the key idea and 1 for a correct explanation or example. At least 4/6 is the lesson target; remediate missed prerequisites.

1. What should a test assertion describe?

2. What does httptest avoid needing?

3. Why isolate database fixtures?

<details>
<summary>Instructor answer key — open after submission</summary>

1. Observable required behavior.

2. A separately deployed HTTP server for handler tests.

3. To avoid order dependence and interference between tests or users.

</details>
