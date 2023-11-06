CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY,
  'role' TEXT NOT NULL,
  'name' TEXT NOT NULL,
  'email' TEXT NOT NULL UNIQUE,
  'password' TEXT NOT NULL,
  'address' TEXT NOT NULL,
  'created_at' DATE NOT NULL,
  'updated_at' DATE NOT NULL
);