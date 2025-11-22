-- This script sets up the PostgreSQL database.

CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(255) NOT NULL UNIQUE,
  username VARCHAR(255) NOT NULL UNIQUE,
  password TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS sessions (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL REFERENCES users(id),
  refresh_token TEXT NOT NULL UNIQUE,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
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
EXECUTE PROCEDURE update_timestamp();

-- Apply the 'update_timestamp' trigger to the sessions table.
CREATE OR REPLACE TRIGGER update_sessions_timestamp
BEFORE UPDATE ON sessions
FOR EACH ROW
EXECUTE PROCEDURE update_timestamp();
