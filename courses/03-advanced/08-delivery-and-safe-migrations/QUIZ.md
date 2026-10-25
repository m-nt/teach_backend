# Quiz: Delivery pipelines, release rollback, and safe migrations

Allow 8 minutes, then 2 minutes for correction. Each question is worth 2 points: 1 for the key idea and 1 for a correct explanation or example. At least 4/6 is the lesson target; remediate missed prerequisites.

1. Does rolling back an image undo a schema migration?

2. Why deploy the same image artifact across environments?

3. When should a destructive schema contraction happen?

<details>
<summary>Instructor answer key — open after submission</summary>

1. No.

2. To reduce differences between what was tested and what runs.

3. After old consumers are retired and compatibility/recovery requirements are met.

</details>
