-- Drop index for email
DROP INDEX IF EXISTS idx_users_email;

-- Drop users table
DROP TABLE IF EXISTS users;

-- Drop enum type for user roles
DROP TYPE IF EXISTS user_role;
