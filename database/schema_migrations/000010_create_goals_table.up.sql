CREATE TABLE IF NOT EXISTS goals (
  id INTEGER PRIMARY KEY,
  'coordinator_id' INTEGER NOT NULL,
  'name' TEXT NOT NULL,
  'description' TEXT,
  'percentage' INTEGER NOT NULL,
  'due_at' DATE NOT NULL,
  'created_at' DATE NOT NULL,
  'updated_at' DATE NOT NULL,

  FOREIGN KEY('coordinator_id') REFERENCES 'users'('id')
);
