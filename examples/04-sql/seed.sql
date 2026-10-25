-- Apply once after schema.sql, to a fresh database only.
BEGIN;
INSERT INTO students (name) VALUES ('Ada'), ('Sam');
INSERT INTO class_sessions (title, capacity) VALUES ('Go fundamentals', 2), ('Go intermediate', 12);
INSERT INTO bookings (student_id, session_id) VALUES (1, 1);
COMMIT;
