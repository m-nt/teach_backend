# Quiz: Context, timeouts, defer, and resource lifetimes

Allow 8 minutes, then 2 minutes for correction. Each question is worth 2 points: 1 for the key idea and 1 for a correct explanation or example. At least 4/6 is the lesson target; remediate missed prerequisites.

1. Does a context forcibly kill any arbitrary function?

2. Why call cancel even after success?

3. When does defer execute?

<details>
<summary>Instructor answer key — open after submission</summary>

1. No, the work must observe cancellation or use context-aware operations.

2. To release context-associated resources promptly.

3. When the surrounding function returns.

</details>
