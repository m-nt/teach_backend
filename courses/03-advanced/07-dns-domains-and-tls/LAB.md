# Lab: DNS, domains, reverse proxies, and TLS

**Starting point:** your prior capstone commit and today’s prerequisite preparation.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Resolve the chosen hostname and verify the expected address.
2. Configure the reverse proxy and keep the API/database behind it.
3. Obtain a lab or public certificate using the selected track.
4. Check hostname validation, HTTP redirection where configured, certificate expiry, and renewal storage.

## Acceptance checks

- [ ] Client verifies the certificate without -k.
- [ ] The chosen hostname reaches the API.
- [ ] Student can isolate a DNS failure from a TLS failure.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
