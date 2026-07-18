INSERT INTO users (first_name, last_name, email, password, is_owner)
VALUES ('Rahim', 'Uddin', 'rahim.uddin@example.com', 'SecurePass123!', false)
RETURNING *;