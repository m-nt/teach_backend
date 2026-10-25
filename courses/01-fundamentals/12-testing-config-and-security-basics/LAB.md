# Lab: Testing, configuration, and security basics

**Starting point:** the referenced repository example or an empty personal exercise folder.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Write a test for the cancellation rule from lesson 04.
2. Run the test, introduce a real rule bug, observe failure, and repair it.
3. Change a local API port through an environment variable where supported.
4. Review a fictional log containing password and OTP fields; rewrite it to preserve useful context without secrets.

## Acceptance checks

- [ ] The new test fails for the deliberate bug and passes after repair.
- [ ] Student distinguishes authentication from authorization.
- [ ] No credentials appear in the submitted evidence.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
