# PostgreSQL student registration API

First start the [local database](../../infrastructure/local/README.md) and apply schema/seed **once**. Then, inside this directory:

```sh
go mod download
export DATABASE_URL='postgres://course:course-local-only@localhost:5432/backend_class?sslmode=disable'
go run .
```

PowerShell: use `$env:DATABASE_URL = "postgres://course:course-local-only@localhost:5432/backend_class?sslmode=disable"` before `go run .`. These are deliberately public training credentials for loopback-only Docker; `sslmode=disable` is for this local exercise, not a public deployment.

In a second terminal:

```sh
curl -i http://localhost:8080/students
curl -i http://localhost:8080/sessions
curl -i -H 'Content-Type: application/json' --data '{"name":"Mina"}' http://localhost:8080/students
curl -i -H 'Content-Type: application/json' --data '{"name":" "}' http://localhost:8080/students
```

Expected: reads 200, valid creation 201 with a new ID, blank name 400. Restart the Go process and read again: Mina remains. Stop `db` briefly and read: 503 with a generic database error. Restart it afterward. Startup refuses to run without a configured reachable database. `ADDR` can select another loopback port.

The example intentionally limits list queries to 100 rows without a full pagination contract. It has no authentication, email uniqueness, durable booking endpoint, or production migration runner. All database failures map to a teaching-friendly 503, including schema mistakes; intermediate improves error classification. The driver is pinned for reproducibility, not claimed to be the latest; review dependency versions and run a vulnerability scan before any real deployment.

`go test ./...` checks input parsing only; it is **not a database integration test**. The manual lifecycle above is the persistence acceptance check. Instructor follow-up exercise: add GET /students/count and test it against the real lab database. Explain the supplied context/shutdown scaffolding at a high level; intermediate teaches those mechanics deeply.
