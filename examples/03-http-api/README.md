# In-memory HTTP API

Run `go run .`, then use a second terminal:

```sh
curl -i http://localhost:8080/sessions
curl -i -H 'Content-Type: application/json' --data '{"session_id":1,"student_name":"Ada"}' http://localhost:8080/bookings
curl -i -H 'Content-Type: application/json' --data '{' http://localhost:8080/bookings
```

`GET /sessions`: 200 with a one-element array. A valid reservation: 201 with booking/session IDs. Bad JSON or invalid fields: 400. Missing session: 404. After two accepted reservations: 409 (full). Use `go test -race ./...` to exercise failure and concurrent-capacity behavior.

Restart to reset all reservations. The mutex is supplied because handlers can run concurrently; students learn it in depth later. Student identity is **not authenticated**; repeated names can reserve more than once. No durable booking record, session tokens, or database is implemented. Do not expose this demo publicly.

Set `ADDR=127.0.0.1:8081 go run .` to avoid a local port conflict in Bash. In PowerShell set `$env:ADDR = "127.0.0.1:8081"` before `go run .`.
