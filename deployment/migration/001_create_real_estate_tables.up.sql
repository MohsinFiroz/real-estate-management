-- Create enum type for user roles
CREATE TYPE user_role AS ENUM ('admin', 'user');

-- Create users table
CREATE TABLE users
(
    id            VARCHAR(26) PRIMARY KEY,
    email         VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    first_name    VARCHAR(50)  NOT NULL,
    last_name     VARCHAR(50),
    phone         VARCHAR(20),
    role          user_role    NOT NULL    DEFAULT 'user',
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    last_login    TIMESTAMP WITH TIME ZONE,
    is_active     BOOLEAN                  DEFAULT true,
);

-- Create index on email for faster lookups
CREATE INDEX idx_users_email ON users (email);