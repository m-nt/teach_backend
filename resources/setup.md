# Student setup and troubleshooting

No previous programming knowledge is assumed. Install tools in stages. A laptop with 8 GB RAM is a practical target for Go plus the small local stack; a full local Kubernetes lab often benefits from 16 GB. These are classroom planning estimates, not vendor minimums. Pairing or an instructor machine is the fallback for limited hardware. No paid cloud subscription is required.

## Before fundamentals lesson 01

1. Install Go using the official Go installation guide in [sources.md](sources.md). Use the same supported stable toolchain selected by the instructor. Examples declare a Go 1.22 language baseline, **not** a recommendation to install an old unsupported release.
2. Install an editor (for example VS Code) and its official Go extension, or use an existing editor. Allow time for extension tools to download.
3. Open a terminal. Run `go version`. If it says command not found, restart the terminal after installation and check the installer’s PATH instructions.
4. Create and reopen a practice folder. On Windows use PowerShell or a supported WSL environment; do not mix a Windows Go installation with Linux paths accidentally.
5. Open `examples/01-hello`, run `go run .`, and save the output. If you cannot finish, bring the exact error and use a supported classroom machine.

## Before lesson 06

Install Git from its official site; verify `git --version`. Set a local author name/email if needed. An online hosting account is optional; local commits are sufficient. Never submit a password or access token as evidence.

## Before lesson 07

Verify `curl --version`. In Windows PowerShell use `curl.exe` when `curl` resolves to another command. Bash examples use single-quoted JSON; for shell differences put JSON in `request.json` and use `curl.exe -i -H "Content-Type: application/json" --data-binary @request.json http://localhost:8080/bookings`.

## Before lesson 09

Install Docker Desktop or Docker Engine plus the Compose plugin using the official OS instructions. Check virtualization availability and start the Docker service/application. Run:

```sh
docker version
docker compose version
```

Client output alone is insufficient: `docker version` must reach the server. On Linux follow the documented permissions setup; do not fix permissions by making the Docker socket world-writable. A blocked corporate laptop needs a shared lab machine, not improvised privilege changes.

During the preceding week, the instructor should pull the images from `infrastructure/local/compose.yaml` and cache them for the classroom. Review applicable software licensing for the venue’s organization as part of tool selection.

## Before lesson 10 / 11

Follow [the local services guide](../infrastructure/local/README.md), verify `SELECT 1`, then apply the SQL scripts. Before lesson 11 run `go mod download` inside `examples/05-db-api` while internet access is available. It is the only supplied Go example needing an external module.

## Before advanced Kubernetes lessons

Install kubectl and either kind or minikube using official instructions. Pick one class-wide tool and record its version. Ensure images can be loaded or pulled. Verify `kubectl cluster-info`. A single shared cluster can be used with per-team namespaces and quotas; do not hand out cluster-admin access as the normal student workflow.

## Common failures

| Symptom | Check first | Safe next step |
|---|---|---|
| `go.mod file not found` | Current folder | Open the example directory, then run `go run .` |
| Program output did not change | Saved file and working directory | Save, confirm the file path, rerun |
| Port already in use | Another API or database instance | Stop your own old process or choose a documented alternate port |
| Docker daemon unreachable | Service running and user access | Start the service or use the instructor machine |
| Cannot pull/download | Internet, proxy, registry access | Use the instructor’s prepared dependency/image cache |
| API cannot reach DB | Host versus container address | Host uses localhost; container uses the Compose service name |
| Authentication failed after changing Compose password | Existing volume still has original credentials | Use the original lab credentials or follow an explicit disposable-data reset; changing environment does not rewrite the existing database user |

## Before every class

Charge the laptop, bring the previous exercise, run the requested preparation command, and record the exact error if it fails. Keep a backup copy of your code. The instructor verifies prerequisites early enough to arrange a fallback.
