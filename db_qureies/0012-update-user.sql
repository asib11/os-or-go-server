-- update user email
UPDATE users
SET email = 'new.email@example.com'
WHERE id = 1
RETURNING *;

-- update many fields at once
UPDATE users
SET first_name = 'Karim', last_name = 'Ahmed', is_owner = true
WHERE id = 1
RETURNING *;