# SQL schema and fixtures

Follow [local service setup](../../infrastructure/local/README.md) to apply these files. Apply schema and seed **once to a fresh database**; they intentionally fail instead of silently modifying an existing schema. Use a separate new database for a repeated clean-start exercise.

The model has students, sessions, and bookings. To keep the first SQL lab small, a session stores its title directly; extracting a reusable Course table is a capstone extension. Foreign keys enforce references and the unique pair prevents a duplicate student/session booking. They do not enforce total class capacity; intermediate lesson 07 adds transactional coordination.

Expected seed results: two students, two sessions, one booking (Ada in Go fundamentals). The remaining-seats query returns 1 and 12. Try duplicate and invalid-reference inserts inside an explicit transaction, then `ROLLBACK`; do not assume a failed statement resets the transaction for you.
