-- This script sets up the PostgreSQL database.

CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(255) NOT NULL UNIQUE,
  username VARCHAR(255) NOT NULL UNIQUE,
  password TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS sessions (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  refresh_token TEXT NOT NULL UNIQUE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

  CONSTRAINT sessions_user_id_fkey
    FOREIGN KEY(user_id) 
    REFERENCES users(id) 
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS wishlists (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  name VARCHAR(255) NOT NULL,
  visibility VARCHAR(20) NOT NULL DEFAULT 'PRIVATE',
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

  CONSTRAINT wishlists_user_id_fkey
    FOREIGN KEY(user_id) 
    REFERENCES users(id) 
    ON DELETE CASCADE,

  CONSTRAINT check_visibility 
    CHECK (visibility IN ('PUBLIC', 'FRIENDS_ONLY', 'PRIVATE'))
);

-- A trigger to automatically update the 'updated_at' column.
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER LANGUAGE PLPGSQL AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$;

-- Apply the 'update_timestamp' trigger to the users table.
CREATE OR REPLACE TRIGGER update_users_timestamp
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

-- Apply the 'update_timestamp' trigger to the sessions table.
CREATE OR REPLACE TRIGGER update_sessions_timestamp
BEFORE UPDATE ON sessions
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

-- Apply the 'update_timestamp' trigger to the wishlists table.
CREATE OR REPLACE TRIGGER update_wishlists_timestamp
BEFORE UPDATE ON wishlists
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();
