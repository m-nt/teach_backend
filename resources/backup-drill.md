# PostgreSQL backup and isolated restore drill

Use synthetic data in the local course stack. Commands run from `infrastructure/local`. This exercise creates a **separate database** called `backend_class_restore`; it never drops or overwrites the source database. If that restore database already exists, choose another unused name and substitute it consistently instead of deleting it blindly.

1. Record a known student ID and current booking count. Record the start time and your recovery targets.
2. Create a custom-format dump (Bash/macOS/Linux):

   ```sh
   docker compose exec -T db pg_dump -U course -d backend_class -Fc > /tmp/backend-class.dump
   ```

   For PowerShell versions that transform redirected binary output, write the dump inside the container with `pg_dump -f /tmp/backend-class.dump`, then copy it using `docker compose cp db:/tmp/backend-class.dump ./backend-class.dump`. Use a writable private output directory on your machine.
3. Add a later synthetic row to the source and note its ID/time. This proves the difference between the backup point and current state.
4. Create a new target and restore:

   ```sh
   docker compose exec -T db createdb -U course backend_class_restore
   docker compose exec -T db pg_restore -U course -d backend_class_restore --no-owner < /tmp/backend-class.dump
   ```

   PowerShell fallback: copy the dump into the container and run `pg_restore ... /tmp/backend-class.dump` without shell input redirection.
5. Query counts, foreign-key relationships, and the known record in the restore target. Confirm the later row is absent if it was created after the dump snapshot.
6. Run a separate API instance against the restore database on an alternate port and exercise a synthetic read/create flow. Do not point the running source application at the target accidentally.
7. Record elapsed restore/verification time and recovered data point. Compare to RTO/RPO targets; list what the lab did not test (off-site loss, encryption/key recovery, production volume).

Backups can contain personal data and secrets. Restrict access and define retention. Delete only the disposable artifacts you intentionally created after the drill; do not include a blanket volume-deletion command in the routine runbook. A successful dump command alone is not a successful restore drill.
