# Kubernetes lab templates

These manifests are scaffolding for the **student’s intermediate/advanced application**, not the beginner examples. They require `/healthz`, `/readyz`, container-visible binding on `0.0.0.0:8080`, and a worker entrypoint. Resource settings are starting guesses to measure, not production sizing.

## Prepare before applying

1. Create a disposable kind/minikube cluster using its official guide. Confirm `kubectl config current-context`; do not use a production cluster.
2. Build API and worker images from the student project. Load them into the local cluster or push to an accessible registry. Replace `REPLACE_WITH_STUDENT_API_IMAGE` and `REPLACE_WITH_STUDENT_WORKER_IMAGE` in a personal copy, preferably with recorded digests. Never assume the cluster can see a host-only image.
3. Create the namespace with `kubectl create namespace backend-class` (skip if already present).
4. Supply reachable PostgreSQL, Redis, and Mailpit addresses. A host Compose service bound to 127.0.0.1 is **not automatically reachable from cluster Pods**. Use isolated in-cluster lab dependencies or explicitly configured private external services. Do not solve this by publicly exposing the lab database.
5. Create `backend-config` and `worker-config` Secrets from locally prepared, gitignored environment files using `kubectl -n backend-class create secret generic backend-config --from-env-file=/path/to/api.env` and the corresponding worker command. Use separate database roles/permissions for API and worker when completing the advanced access-control lab. Example variable names: `DATABASE_URL`, `REDIS_ADDR`, `SMTP_ADDR`; use the actual names your application expects. Secrets require cluster access protection; encoding is not encryption.
6. Apply the prepared API manifest, then the worker manifest. The template namespace object may be reapplied safely.

## Inspect and exercise

```sh
kubectl -n backend-class get deployments,pods,services
kubectl -n backend-class rollout status deployment/api
kubectl -n backend-class port-forward service/api 8080:80
```

Use another terminal for requests. Delete a selected API Pod by its actual name and watch replacement. After replacement, restart port-forward if its selected Pod connection closed; port-forward is a lab tunnel, not a production ingress. Inspect failures with `describe`, `logs`, and events. A working Deployment does not prove DNS/TLS; follow the separate proxy lesson.

## Student implementation tasks

- Add a startup probe if measured initialization needs one.
- Choose and test a network plugin that enforces NetworkPolicy. Write policies from the actual labels and dependency locations, permit DNS as required, and demonstrate both allowed and denied flows. A generic policy without a known topology could block all legitimate work, so none is silently applied here.
- Keep the service accounts unprivileged and API-token automount off unless the application actually needs the Kubernetes API. Add only the reviewed RBAC permissions needed for a concrete operation.
- Implement worker shutdown and observe queue age/worker progress; no fake HTTP health endpoint is assumed for the worker.
- Document database/Redis/mail provisioning, persistent storage, backups, and TLS routing in the student runbook. These are capstone deliverables, not services automatically installed by these manifests.

Cleanup: confirm your context and delete only the disposable `backend-class` namespace when you deliberately finish the lab. Check ownership of persistent volumes first; namespace deletion can remove associated data depending on storage configuration.
