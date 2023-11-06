CREATE TABLE IF NOT EXISTS events (
  id INTEGER PRIMARY KEY,
  'type' INTEGER NOT NULL,
  'description' TEXT NOT NULL,
  'start_at' DATE NOT NULL,
  'end_at' DATE NOT NULL,
  'created_at' DATE NOT NULL,
  'updated_at' DATE NOT NULL
);
