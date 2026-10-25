# Quiz: Schema migrations, indexes, and query plans

Allow 8 minutes, then 2 minutes for correction. Each question is worth 2 points: 1 for the key idea and 1 for a correct explanation or example. At least 4/6 is the lesson target; remediate missed prerequisites.

1. Does every column need an index?

2. Does EXPLAIN ANALYZE execute its query?

3. Why test migrations with old data?

<details>
<summary>Instructor answer key — open after submission</summary>

1. No; indexes should serve measured access patterns.

2. Yes.

3. Constraints and data transformations can fail despite a clean install working.

</details>
