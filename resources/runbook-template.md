# Student operational handoff template

Fill every section using the actual deployed project. Use placeholders for secrets and keep secret values in the selected controlled store.

## Ownership and inventory

Owner/contact: __. Environment and purpose: __. Source revision: __. Image digests: __. Go/database/Redis/Kubernetes versions: __. API hostname: __. Namespace: __. Stateful-service location: __. Port/access matrix: __.

## Deploy and verify

Prerequisites: __. Configuration names and injection method: __. Migration owner and command: __. Deploy command: __. Readiness check: __. Verified TLS request: __. Synthetic booking/notification check: __. Expected output: __.

## Rollback

Trigger/decision owner: __. Previous artifact: __. Schema compatibility: __. Rollback command: __. Verification: __. Irreversible changes and recovery alternative: __.

## Dependency incidents

For database, Redis, notification worker, mail provider, proxy/DNS: symptom, first observations, expected degraded behavior, mitigation command, recovery check, escalation owner. Never put passwords or token-bearing commands into the shared document.

## Data recovery

Backup schedule/location/access: __. RPO/RTO targets: __. Last tested restore and evidence: __. Isolated restore command: __. Business verification: __. Promotion procedure after review: __.

## Security maintenance

Credential rotation procedure: __. Session invalidation: __. OTP abuse review: __. Dependency scan: __. Certificate renewal observation: __. Access review: __.

## Known limitations and next work

List each limitation with the affected behavior, evidence, owner, and a bounded next step. Do not describe an untested control as working.
