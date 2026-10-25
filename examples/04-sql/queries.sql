SELECT id, title, capacity FROM class_sessions ORDER BY id;
SELECT id, name FROM students WHERE id = 1;
SELECT b.id, s.name, c.title
FROM bookings AS b
JOIN students AS s ON s.id = b.student_id
JOIN class_sessions AS c ON c.id = b.session_id
ORDER BY b.id;
SELECT c.id, c.title, c.capacity - COUNT(b.id) AS remaining
FROM class_sessions AS c
LEFT JOIN bookings AS b ON b.session_id = c.id
GROUP BY c.id, c.title, c.capacity
ORDER BY c.id;
