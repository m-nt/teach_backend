# Quiz: Go value semantics, pointers, and methods

Allow 8 minutes, then 2 minutes for correction. Each question is worth 2 points: 1 for the key idea and 1 for a correct explanation or example. At least 4/6 is the lesson target; remediate missed prerequisites.

1. Does assigning a struct create an alias to the original struct?

2. Does copying a slice guarantee independent elements?

3. When is a pointer receiver appropriate?

<details>
<summary>Instructor answer key — open after submission</summary>

1. No, it copies the struct value.

2. No, it can share the backing array.

3. When the method must mutate its receiver or receiver consistency requires it.

</details>
