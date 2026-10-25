# Quiz: DNS, domains, reverse proxies, and TLS

Allow 8 minutes, then 2 minutes for correction. Each question is worth 2 points: 1 for the key idea and 1 for a correct explanation or example. At least 4/6 is the lesson target; remediate missed prerequisites.

1. Does a working DNS lookup prove HTTPS works?

2. Why persist the proxy certificate state?

3. Is -k an acceptable verification check?

<details>
<summary>Instructor answer key — open after submission</summary>

1. No; transport, certificate validation, and routing can still fail.

2. To preserve keys and issuance/renewal state across restarts.

3. No, it bypasses certificate verification.

</details>
