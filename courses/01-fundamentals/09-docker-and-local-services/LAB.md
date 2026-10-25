# Lab: Docker, images, containers, ports, and volumes

**Starting point:** the referenced repository example or an empty personal exercise folder.
**Deliverable:** a saved exercise or commit, observed output, and a short explanation of one success and one failure. Use only synthetic data.

## Steps

1. Read compose.yaml and identify image, ports, environment, health check, and volume.
2. Start only db and inspect its status.
3. Run SELECT 1 through psql inside the container.
4. Stop and restart db. Explain why removing volumes would be a different operation; do not remove them during this exercise.

## Acceptance checks

- [ ] SELECT 1 returns one row.
- [ ] Student identifies host port 5432 and container port 5432.
- [ ] Student explains where database files persist.

## Evidence worksheet

| Case / input | Expected result before running | Actual result | Explanation or repair |
|---|---|---|---|
| Main successful path | Fill before running | Record output | Explain |
| Boundary or failure path | Fill before running | Record output | Explain |

## Pair review

Have your partner reproduce one check without your help. Explain the request or state transition out loud. Save your work before attempting an extension. Stop only the services you started; preserve database volumes unless a separate disposable-data reset is explicitly intended.
