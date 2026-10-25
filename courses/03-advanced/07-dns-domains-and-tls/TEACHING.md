# Instructor notes: DNS, domains, reverse proxies, and TLS

## Before class

Read the student lesson and lab. Rehearse the demonstration on the classroom machine, prepare a known-working copy of the previous lesson’s student project, and retain the failure output. Verify the prerequisite handoff from the previous homework. Use synthetic data. Assign driver/navigator pairs and swap midway.

## Explain and draw

1. DNS maps a hostname to records, such as A/AAAA addresses. DNS resolution, TCP reachability, TLS validation, and HTTP routing are separate layers; troubleshoot them in that order.

2. A reverse proxy accepts public traffic and forwards to a private API. TLS protects traffic and authenticates the server name when the client trusts the issuing authority. “SSL” is common shorthand; teach modern TLS terminology.

3. Public automatic certificates require control of the name and a successful validation challenge. For offline/local practice use Caddy internal TLS and explicitly trust the local root on the lab client. Do not teach disabling certificate verification as the solution.

## Demonstration script

Use infrastructure/Caddyfile.local with a local API, explain the internal CA, and verify with curl --cacert after exporting the lab root. For a public demo use a controlled DNS name and Caddy automatic HTTPS.

Before each command or code change, ask students to predict the result. Afterward, ask which observation supports their prediction. Keep the initial working version available so one installation problem does not consume the whole lab.

## Worked reasoning / expected solution

- Does a working DNS lookup prove HTTPS works? **Expected reasoning:** No; transport, certificate validation, and routing can still fail.
- Why persist the proxy certificate state? **Expected reasoning:** To preserve keys and issuance/renewal state across restarts.
- Is -k an acceptable verification check? **Expected reasoning:** No, it bypasses certificate verification.

## Review the lab

- Client verifies the certificate without -k.
- The chosen hostname reaches the API.
- Student can isolate a DNS failure from a TLS failure.

## If learners get stuck

First reproduce the exact failure and read the first relevant error. Check the working directory, saved files, running dependencies, and previous prerequisite before changing code. Pair a blocked student with a working machine; the blocked student should explain and perform the next change. If setup remains broken, use the demonstrated request/state diagram and record the repair as a clinic task.

For learners who finish early: add one boundary or failure case directly related to today’s outcome and explain why it matters. Do not require optional extensions to pass.

## End-of-class record

Record each learner’s exit-check status (independent / with help / needs repair), the exact missing behavior, and the next preparation task. Discuss quiz answers after students submit; answers are teaching aids, not a secret examination bank.
