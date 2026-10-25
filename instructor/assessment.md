# Assessment, placement, and feedback

Assess working behavior and understanding. Students may consult documentation; ask them to explain their own changes. Allow corrections after feedback. Do not award points simply for using more tools.

## Per-lesson assessment

Each quiz has 3 questions × 2 points. Award 1 for the core idea and 1 for a correct explanation/example. Target 4/6. Each homework is also 6 points: observable behavior (2), explanation (2), reproducible evidence (2). Missing preparation triggers setup support, not automatic exclusion. Require repair before a dependent task if the missing prerequisite prevents participation.

Record: learner, lesson, exit check (independent / with help / needs repair), quiz, homework, precise blocker, repair deadline. Use synthetic identifiers in shared records; do not publish student scores in the repository.

## Tier grade

- Homework and lab portfolio: 30%.
- Quizzes and explanation checks: 20%.
- Capstone rubric: 50%.

Pass at 70/100 overall **and** pass all mandatory tier gates. A serious correctness/security gap cannot be averaged away with quiz points. Offer one-week repair/resubmission and a short follow-up explanation. Attendance is a participation policy set by the organizer, not evidence of competence.

## Capstone rubric (100 points)

| Dimension | Points | Full-credit evidence |
|---|---|---|
| Required behavior | 30 | All tier acceptance flows reproduce |
| Correctness and failure handling | 25 | Boundary, dependency, and applicable concurrency cases pass |
| Individual understanding | 20 | Learner traces a request, explains ownership, and makes a small change |
| Reproducibility | 15 | Peer can run/deploy using the guide and versioned artifacts |
| Clarity and limitations | 10 | Domain diagram, useful naming, honest limitations, redacted evidence |

Within each dimension, award full points for independent and reproducible work, about half for partially working work needing substantial guidance, and zero for absent evidence. Explain deductions with a concrete failed behavior.

## Mandatory gates

**Fundamentals:** run/edit a Go program; explain functions and a business rule; handle invalid input; demonstrate a SQL relationship and persistence across API restart; execute at least one meaningful rule test; avoid committed secrets. The booking demo may remain in memory, but registration must persist.

**Intermediate:** all above plus no overbooking under simultaneous requests, cross-user resource access denied, safe password/session handling, cache-down fallback, reproducible tests and container startup. A failed authorization or capacity invariant requires repair regardless of total points.

**Advanced:** all above plus durable delivery intent and recoverable worker claims, bounded OTP issuance/attempts and single-use verification, verified TLS without bypass, Kubernetes rollout/Pod recovery evidence, tested restore to an isolated database, and an application rollback compatible with the schema. NetworkPolicy enforcement must be tested on a supporting environment to claim that outcome.

## Placement without taking an earlier tier

For intermediate, give 45–60 minutes: run a small Go project, repair a rule, explain an HTTP error, write a SQL join, and show a basic test. Require every fundamentals gate; assign a bridge exercise for gaps.

For advanced, request an existing project and a 60–90-minute practical review: persistent concurrent booking, two-user authorization test, container startup, and dependency-failure handling. Do not replace this with a vocabulary quiz.

## Feedback template

“Observed: __. Expected: __. Reproduce with: __. Next repair: __. Evidence to resubmit: __.” Praise specific reasoning or debugging behavior, and separate installation difficulties from programming understanding.
