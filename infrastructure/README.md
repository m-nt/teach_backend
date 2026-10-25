# Infrastructure teaching tracks

[Local dependencies](local/README.md) support fundamentals and intermediate. They do not include a completed student API or worker.

## Local TLS without a purchased domain

Install Caddy from its official documentation and run from this directory:

```sh
caddy run --config Caddyfile.local
```

Run the API separately at `127.0.0.1:8080`. Caddy serves `https://localhost:8443` using its internal CA. Find the data directory for your OS using Caddy’s documentation, export the local CA root certificate (not the private key), and verify:

```sh
curl --cacert /path/to/lab-root.crt https://localhost:8443/sessions
```

The root certificate path is intentionally environment-specific. Verify its identity on the instructor machine before trusting it. Importing a root CA changes a machine’s trust settings; using `--cacert` for the single lab command keeps trust scoped. Do not use `-k` as proof of TLS verification. With this Caddyfile there is no assumed separate HTTP listener or redirect test.

For the optional public track, use a controlled hostname as the Caddy site address, remove `tls internal`, configure the required DNS and validation reachability, and retain Caddy’s data directory. Follow the official automatic-HTTPS guidance. Do not expose the local database or reuse the lab’s public credentials. Test actual issuance and renewal behavior before claiming that outcome.

## Kubernetes track

The [Kubernetes template guide](k8s/README.md) explains what must be supplied and what has not been automated. API and worker are student artifacts; manifests deliberately require image/configuration replacement. The [Dockerfile example](../resources/Dockerfile.example) is a starting point for a student module with `cmd/api`; use a similar build target for `cmd/worker`.
