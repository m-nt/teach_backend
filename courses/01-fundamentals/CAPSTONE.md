# Fundamentals capstone: local class directory and registration

Build a local Go application that lists class sessions and persists student registration. Keep booking rules as a tested CLI or an explicitly in-memory demo. This is the 35-hour beginner outcome; durable concurrent booking belongs to intermediate.

## Starting material and required student work

Use [the examples](../../examples/README.md) as scaffolding. Combine the database API with a session-list endpoint, write at least one new validation rule and its test, and write your own run guide and domain diagram. Explain any copied scaffold. Do not submit the unmodified examples as a finished capstone.

## Mandatory acceptance flow

1. A peer follows your README to start PostgreSQL and the Go process.
2. `GET /sessions` returns the seeded session catalog with stable IDs.
3. `POST /students` creates a synthetic student and returns 201.
4. A blank or whitespace-only name returns 400 without inserting a row.
5. `GET /students` returns the student after restarting the Go process.
6. Show a SQL join relating students, sessions, and the supplied booking fixture.
7. Run your booking-rule tests: success, full class, and invalid input. Show one test detects a deliberate bug, then restore the working implementation.
8. Stop the database briefly; a request fails with a controlled error rather than crashing the process. Restore the database.

## Submission

- Source and a commit ID; no credentials beyond documented local-only training defaults.
- Entity drawing with Student, Course, ClassSession, and Booking; explain the simplified starter schema where a session stores its title directly.
- Startup and verification commands, test output, and one error-debugging story.
- Known limitations: local-only, no authentication, synthetic records, and any in-memory state.

Use the [assessment rubric](../../instructor/assessment.md). Optional extension: persist a course table and join it to sessions. Do not add authentication or cloud infrastructure to compensate for missing core behavior.
