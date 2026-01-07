
INSERT INTO users (name, dob, email, password_hash, role) 
VALUES ($1, $2, $3, $4, COALESCE($5, 'user')) 
RETURNING id, name, dob, email, role, created_at, updated_at;


SELECT id, name, dob, email, role, created_at, updated_at 
FROM users 
WHERE id = $1;

SELECT id, name, dob, email, password_hash, role, created_at, updated_at 
FROM users 
WHERE email = $1;

SELECT id, name, dob, email, role, created_at, updated_at 
FROM users 
ORDER BY id;

SELECT id, name, dob, email, role, created_at, updated_at 
FROM users 
ORDER BY id
LIMIT $1 OFFSET $2;

SELECT COUNT(*) 
FROM users;

UPDATE users 
SET name = $2, dob = $3, updated_at = CURRENT_TIMESTAMP 
WHERE id = $1 
RETURNING id, name, dob, email, role, created_at, updated_at;

UPDATE users 
SET password_hash = $2, updated_at = CURRENT_TIMESTAMP 
WHERE id = $1 
RETURNING id, email, updated_at;

DELETE FROM users 
WHERE id = $1;

