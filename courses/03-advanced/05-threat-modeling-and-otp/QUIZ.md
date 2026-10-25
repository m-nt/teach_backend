# Quiz: Threat modeling, OTP, and abuse controls

Allow 8 minutes, then 2 minutes for correction. Each question is worth 2 points: 1 for the key idea and 1 for a correct explanation or example. At least 4/6 is the lesson target; remediate missed prerequisites.

1. Why is plain hashing of a short numeric code insufficient against offline guessing?

2. Why consume atomically?

3. Is email OTP automatically phishing-resistant MFA?

<details>
<summary>Instructor answer key — open after submission</summary>

1. The small code space is cheap to enumerate; a keyed digest adds protection when the store alone is exposed.

2. Concurrent verifications must not both accept the same challenge.

3. No.

</details>
