# Roadmaps and primary references

Curriculum prepared on **2026-10-25**. The backend/Go roadmaps and core official documentation landing pages were consulted for sequencing and scope. Topic links below are instructor reading references, not a guarantee that every version-specific page or installation command was exercised. Recheck versions and security recommendations before each cohort. The OWASP password-storage page returned HTTP 403 during verification; its current contents were not confirmed in this environment. See [the validation record](../instructor/validation.md). Lessons are original explanations and exercises, not copied roadmap content.

| Topic / lessons | Reference | How to use it |
|---|---|---|
| Overall breadth | https://roadmap.sh/backend | Check the broad order: language, HTTP, relational data, auth, cache, testing, delivery |
| Go progression | https://roadmap.sh/golang | Check language/tooling topics; do not assign the entire map to beginners |
| Go basics, packages, database, tests | https://go.dev/doc/ and https://go.dev/doc/tutorial/ | Rehearse the relevant official tutorial before teaching |
| Go installation | https://go.dev/doc/install | Follow the learner’s actual OS instructions |
| Standard library | https://pkg.go.dev/std | Look up net/http, context, database/sql, testing, crypto/rand |
| Race detection | https://go.dev/doc/articles/race_detector | Explain capabilities and limits before the concurrency lab |
| SQL | https://www.postgresql.org/docs/current/tutorial.html | Tables, joins, constraints, transactions |
| PostgreSQL operations | https://www.postgresql.org/docs/current/backup.html | Backup/restore choices and operational caveats |
| PostgreSQL driver used by example | https://github.com/jackc/pgx | pgx database/sql adapter and release compatibility |
| Docker / Compose | https://docs.docker.com/get-started/ and https://docs.docker.com/compose/ | OS setup, images, services, ports, volumes |
| Git | https://git-scm.com/book/en/v2 | Local workflow and reviewing changes |
| Editor Go support | https://code.visualstudio.com/docs/languages/go | Optional editor configuration |
| HTTP semantics | https://www.rfc-editor.org/rfc/rfc9110 | Intermediate contract decisions; select relevant sections |
| API description | https://spec.openapis.org/oas/latest.html | Choose and record a concrete schema version for the class |
| Authentication | https://cheatsheetseries.owasp.org/cheatsheets/Authentication_Cheat_Sheet.html | Threats, responses, and abuse controls |
| Password storage | https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html | Choose maintained algorithms/parameters; do not invent cryptography |
| Sessions and CSRF | https://cheatsheetseries.owasp.org/cheatsheets/Session_Management_Cheat_Sheet.html and https://cheatsheetseries.owasp.org/cheatsheets/Cross-Site_Request_Forgery_Prevention_Cheat_Sheet.html | Browser-session design and failure tests |
| Recovery / verification tokens | https://cheatsheetseries.owasp.org/cheatsheets/Forgot_Password_Cheat_Sheet.html | Expiry, single use, enumeration, and request limits |
| Redis | https://redis.io/docs/latest/develop/ | Commands, expiration, transactions/scripts, client behavior |
| Mail sandbox | https://mailpit.axllent.org/docs/install/docker/ | Local-only SMTP/inbox configuration |
| HTTPS | https://caddyserver.com/docs/automatic-https | Public versus internal CA trust and certificate storage |
| Kubernetes | https://kubernetes.io/docs/tutorials/kubernetes-basics/ | Deploy, expose, scale, update, and inspect workloads |
| Local Kubernetes | https://kind.sigs.k8s.io/docs/user/quick-start/ and https://minikube.sigs.k8s.io/docs/start/ | Select one supported local lab environment |
| Kubernetes security | https://kubernetes.io/docs/concepts/security/ | RBAC, workload and network boundaries |
| Telemetry | https://opentelemetry.io/docs/languages/go/ and https://prometheus.io/docs/practices/naming/ | Trace propagation and bounded metric labels |
| Go vulnerability scanning | https://go.dev/doc/security/ | Add vulnerability checks to the quality gate |

Record the exact Go, Docker, database, Redis, Kubernetes, driver, and image versions actually used in the cohort run sheet. Roadmaps are navigation aids; primary documentation is the technical authority for each selected tool.
