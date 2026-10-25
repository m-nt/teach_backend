# Quiz: Packages, small interfaces, and error contracts

Allow 8 minutes, then 2 minutes for correction. Each question is worth 2 points: 1 for the key idea and 1 for a correct explanation or example. At least 4/6 is the lesson target; remediate missed prerequisites.

1. Must a Go type declare that it implements an interface?

2. Why use %w when wrapping?

3. Where should HTTP status selection live?

<details>
<summary>Instructor answer key — open after submission</summary>

1. No, matching methods satisfy it implicitly.

2. It preserves the error chain for inspection.

3. At the transport boundary, not inside the domain rule.

</details>
