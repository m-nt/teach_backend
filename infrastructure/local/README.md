# Local classroom dependencies

These services are for synthetic data on the classroom computer. Published ports bind to loopback. Credentials are public lab defaults. Redis/Mailpit are intentionally unauthenticated inside this local stack; do not publish them to the internet. Image tags identify teaching baselines and may move within a line; record resolved digests for a cohort and review updates before real deployment.

From this directory:

```sh
docker compose config -q
docker compose up -d db
docker compose ps
docker compose exec db psql -U course -d backend_class -c 'SELECT 1;'
```

Apply schema and fixture once to a **fresh** database (Bash/macOS/Linux):

```sh
docker compose exec -T db psql -v ON_ERROR_STOP=1 -U course -d backend_class < ../../examples/04-sql/schema.sql
docker compose exec -T db psql -v ON_ERROR_STOP=1 -U course -d backend_class < ../../examples/04-sql/seed.sql
docker compose exec -T db psql -v ON_ERROR_STOP=1 -U course -d backend_class < ../../examples/04-sql/queries.sql
```

PowerShell alternative, one file at a time:

```powershell
Get-Content -Raw ../../examples/04-sql/schema.sql | docker compose exec -T db psql -v ON_ERROR_STOP=1 -U course -d backend_class
Get-Content -Raw ../../examples/04-sql/seed.sql | docker compose exec -T db psql -v ON_ERROR_STOP=1 -U course -d backend_class
Get-Content -Raw ../../examples/04-sql/queries.sql | docker compose exec -T db psql -v ON_ERROR_STOP=1 -U course -d backend_class
```

The schema/seed is intentionally not rerunnable on an existing database. For a new exercise, create a distinctly named empty database with `createdb` and apply the scripts there. Do not erase the existing volume just to clear an error.

Start later dependencies only when needed:

```sh
docker compose up -d redis mailpit
docker compose exec redis redis-cli PING
docker compose logs --tail=50 db redis mailpit
```

Mailpit inbox: `http://localhost:8025`. SMTP from a host process: `localhost:1025`; from an application on the same Compose network: `mailpit:1025`. Host database address: `localhost:5432`; network service address: `db:5432`. Redis follows the same pattern on 6379. The application is **not** included here; intermediate lesson 12 adds the student-built image to their own Compose setup.

Stop services with `docker compose stop`; `docker compose down` removes containers/network while retaining the named database volume. Do not add `-v` unless deliberately deleting your disposable database. Pulls need network access; instructor image caching should happen before class. Startup/health checks must be rehearsed locally; a valid Compose file alone does not prove the services run.
