-- all user data
SELECT * FROM users;

-- specific columns
SELECT id, first_name, last_name, email FROM users;

-- filtering by condition
SELECT * FROM users WHERE email = 'rahim.uddin@example.com';

-- filtering by multiple conditions
SELECT * FROM users WHERE id = 1;