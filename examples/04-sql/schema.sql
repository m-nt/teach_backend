-- Apply once to a fresh training database. Numbered migrations come later.
BEGIN;
CREATE TABLE students (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL CHECK (length(trim(name)) BETWEEN 1 AND 100)
);
CREATE TABLE class_sessions (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title TEXT NOT NULL,
    capacity INTEGER NOT NULL CHECK (capacity > 0)
);
CREATE TABLE bookings (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    student_id BIGINT NOT NULL REFERENCES students(id),
    session_id BIGINT NOT NULL REFERENCES class_sessions(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (student_id, session_id)
);
COMMIT;
-- These constraints do NOT enforce the cross-row capacity invariant.
-- Intermediate adds a transaction that locks the class session before checking capacity.
