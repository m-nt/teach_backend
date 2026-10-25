# Validation record

Validated on **2026-10-25**. Source/configuration checks are distinguished from runtime demonstrations. These results validate the supplied teaching examples, not the student implementations described by the later capstones.

## Executed successfully

| Check | Result |
|---|---|
| Curriculum structure | 3 tiers, 42 sequenced lessons, 210 required lesson files |
| Local Markdown targets | All links resolve using `python3 scripts/check-curriculum.py` |
| Go formatting, vet, build | All four supplied modules pass |
| Go unit/race checks | CLI rules, HTTP validation/concurrent capacity, database-input parsing pass; greeting has no tests |
| Compose syntax | `docker compose -f infrastructure/local/compose.yaml config -q` passes |
| PostgreSQL schema/seed | Applied to a fresh, isolated PostgreSQL 17.11 container |
| SQL queries | Joins and remaining-seat query return expected seeded records |
| Database HTTP flow | Session/student reads 200; student creation 201; blank name 400 with no inserted row |
| Persistence | Created student remains after API restart |
| Dependency failure | Pausing the isolated database produces 503; reads recover after resume |
| Backup/restore | Custom-format dump restored into a separate database; API reads confirm recovered student and sessions |
| Cleanup | Temporary API process stopped and disposable database container removed |

The database runtime check used loopback port 15432 and ephemeral storage under the uniquely named `teach-backend-validation-20261025` container. Existing user services and their data were not used. The test exercised the documented dump/restore operations, not a timed production recovery or off-site backup policy.

Environment: Go **1.27.0 linux/amd64**, Compose **v5.1.4**, Docker Engine **29.1.3**, PostgreSQL **17.11** from `postgres:17-alpine` (observed digest `sha256:eee6c6e1b0894fa4fe0075eb054af9fd1cf39a35321c2d4110626cfe506411fb`). Example modules use a Go 1.22 language baseline and the database module pins pgx **v5.7.2** with a checked-in checksum file. Compilation on this toolchain does not verify every supported OS or older toolchain.

## Not executed / instructor rehearsal still required

- Full Compose lifecycle with Redis and Mailpit; only the Compose configuration and an isolated PostgreSQL runtime were checked.
- Image build from the student Dockerfile template, since it expects a future student module with `cmd/api`.
- Kubernetes apply/rollout, network-policy enforcement, public DNS, local/public certificate issuance, and certificate renewal.
- Intermediate/advanced student implementations: password sessions, CSRF, transaction-safe persistent reservation, cache invalidation, outbox, OTP, retry/backpressure, and telemetry. Their tests and acceptance criteria are assignments, not claimed completed application features.
- Vulnerability scanning and confirmation that the pinned dependency versions are the newest. An optional driver-version lookup was not executed because automatic approval review hit its usage limit; the existing tested pin was retained.
- A real cohort timing study. Contact-hour plans and homework estimates are instructional estimates; adjust after exit-check evidence.

## References checked

The public backend/Go roadmap pages and Go/Kubernetes documentation landing pages were read to check scope and reference targets. The OWASP password-storage page returned HTTP 403 in this environment; it remains a primary reference to review before teaching authentication. The remaining topic links are reading references, not a claim that every page or version-specific command was exercised.

## Reproduce before class

Run `bash scripts/check-examples.sh` and `python3 scripts/check-curriculum.py` from the repository root. The Go check script downloads external modules into a temporary module cache if missing. Follow the local-service README and database-example README for persistence checks, and `resources/backup-drill.md` for an isolated restore. Rehearse the actual classroom OS, network access, ports, and credentials before learners arrive.
