-- This script sets up the PostgreSQL database.

-- Create the `users` table.
CREATE TABLE IF NOT EXISTS users (
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  email VARCHAR(255) NOT NULL UNIQUE,
  username VARCHAR(36) NOT NULL UNIQUE,
  password TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Create the `sessions` table.
CREATE TABLE IF NOT EXISTS sessions (
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  user_id BIGINT NOT NULL,
  refresh_token_id TEXT NOT NULL UNIQUE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

  -- Add a foreign key reference to the `users.id` column.
  CONSTRAINT sessions_user_id_fkey
    FOREIGN KEY(user_id) 
    REFERENCES users(id) 
    ON DELETE CASCADE
);

-- Create an index for the `sessions.user_id` foreign key.
CREATE INDEX IF NOT EXISTS idx_sessions_user_id_fkey ON sessions(user_id);

-- Create the `wish_lists` table.
CREATE TABLE IF NOT EXISTS wish_lists (
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  user_id BIGINT NOT NULL,
  name VARCHAR(255) NOT NULL,
  visibility VARCHAR(20) NOT NULL DEFAULT 'private',
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

  -- Add a foreign key reference to the `users.id` column.
  CONSTRAINT wish_lists_user_id_fkey
    FOREIGN KEY(user_id) 
    REFERENCES users(id) 
    ON DELETE CASCADE,

  -- Define allowed values for the `visibility` column.
  CONSTRAINT wish_lists_visibility_check
    CHECK (visibility IN ('public', 'friends_only', 'private'))
);

-- Create an index for the `wish_lists.user_id` foreign key.
CREATE INDEX IF NOT EXISTS idx_wish_lists_user_id_fkey ON wish_lists(user_id);

-- Create a trigger to automatically update the `updated_at` column.
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER LANGUAGE PLPGSQL AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$;

-- Apply the `update_timestamp` trigger to the `users` table.
CREATE OR REPLACE TRIGGER update_users_timestamp
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

-- Apply the `update_timestamp` trigger to the `sessions` table.
CREATE OR REPLACE TRIGGER update_sessions_timestamp
BEFORE UPDATE ON sessions
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

-- Apply the `update_timestamp` trigger to the `wish_lists` table.
CREATE OR REPLACE TRIGGER update_wish_lists_timestamp
BEFORE UPDATE ON wish_lists
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();
