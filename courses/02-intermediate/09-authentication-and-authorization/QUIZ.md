# Quiz: Password storage, sessions, and resource authorization

Allow 8 minutes, then 2 minutes for correction. Each question is worth 2 points: 1 for the key idea and 1 for a correct explanation or example. At least 4/6 is the lesson target; remediate missed prerequisites.

1. Why not store passwords using plain SHA-256?

2. Does login permit access to every resource?

3. Why does cookie authentication need CSRF consideration?

<details>
<summary>Instructor answer key — open after submission</summary>

1. It is a fast general hash, not a password-specific work-factor scheme.

2. No, authorization is checked separately.

3. Browsers can attach cookies automatically to requests initiated from another origin.

</details>
