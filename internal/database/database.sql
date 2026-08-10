-- This script sets up the PostgreSQL database.

-- This function automatically sets the `updated_at` column to the current
-- timestamp when a row is updated.
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD IS DISTINCT FROM NEW THEN
    NEW.updated_at = NOW();
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE PLPGSQL;

-------------------------------
-- Set up the `users` table. --
-------------------------------

-- Create the `users` table.
CREATE TABLE IF NOT EXISTS users (
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  email VARCHAR(255) NOT NULL UNIQUE,
  username VARCHAR(36) NOT NULL UNIQUE,
  password TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Apply the `update_timestamp` trigger to the `users` table.
CREATE OR REPLACE TRIGGER users_update_timestamp
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

----------------------------------
-- Set up the `sessions` table. --
----------------------------------

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

-- Create an index for the `sessions.user_id` column.
CREATE INDEX IF NOT EXISTS idx_sessions_user_id ON sessions(user_id);

-- Apply the `update_timestamp` trigger to the `sessions` table.
CREATE OR REPLACE TRIGGER sessions_update_timestamp
BEFORE UPDATE ON sessions
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

------------------------------------
-- Set up the `wish_lists` table. --
------------------------------------

-- Create an enum type for the `wish_lists.visibility` column.
CREATE TYPE wish_lists_visibility_type AS ENUM (
  'public', 'friends_only', 'private'
);

-- Create the `wish_lists` table.
CREATE TABLE IF NOT EXISTS wish_lists (
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  user_id BIGINT NOT NULL,
  name VARCHAR(255) NOT NULL,
  visibility wish_lists_visibility_type NOT NULL DEFAULT 'private',
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

  -- Add a foreign key reference to the `users.id` column.
  CONSTRAINT wish_lists_user_id_fkey
    FOREIGN KEY(user_id) 
    REFERENCES users(id) 
    ON DELETE CASCADE
);

-- Create an index for the `wish_lists.user_id` column.
CREATE INDEX IF NOT EXISTS idx_wish_lists_user_id ON wish_lists(user_id);

-- Apply the `update_timestamp` trigger to the `wish_lists` table.
CREATE OR REPLACE TRIGGER wish_lists_update_timestamp
BEFORE UPDATE ON wish_lists
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

-------------------------------
-- Set up the `items` table. --
-------------------------------

-- Create an enum type for the `items.status` column.
CREATE TYPE items_status_type AS ENUM (
  'available', 'reserved', 'received'
);

-- Create the `items` table.
CREATE TABLE IF NOT EXISTS items (
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  wish_list_id BIGINT NOT NULL,
  url TEXT NOT NULL,
  name VARCHAR(255) NOT NULL,
  price NUMERIC(12, 2) NOT NULL,
  status items_status_type NOT NULL DEFAULT 'available',
  reserved_by_user_id BIGINT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

  -- Add a foreign key reference to the `wish_lists.id` column.
  CONSTRAINT items_wish_lists_id_fkey
    FOREIGN KEY(wish_list_id) 
    REFERENCES wish_lists(id) 
    ON DELETE CASCADE
  
  -- Add a foreign key reference to the `users.id` column.
  CONSTRAINT items_reserved_by_user_id_fkey
    FOREIGN KEY(reserved_by_user_id) 
    REFERENCES users(id)
    ON DELETE SET NULL
);

-- Create an index for the `items.wish_list_id` column.
CREATE INDEX IF NOT EXISTS idx_items_wish_lists_id ON items(wish_list_id);

-- Apply the `update_timestamp` trigger to the `items` table.
CREATE OR REPLACE TRIGGER items_update_timestamp
BEFORE UPDATE ON items
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();
