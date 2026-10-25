# Lab: Git, debugging, and a small project structure

**Starting point:** the referenced repository example or an empty personal exercise folder.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Initialize a repository and configure your local author identity.
2. Add a .gitignore for .env, binaries, and local dumps.
3. Move the booking rule into booking.go in the same package and run go run .
4. Commit a working version; change a rule, inspect the diff, repair it, and commit again.

## Acceptance checks

- [ ] git log shows two meaningful commits.
- [ ] go run . works with multiple source files; no secrets are tracked.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
