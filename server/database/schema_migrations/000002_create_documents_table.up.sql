CREATE TABLE IF NOT EXISTS documents (
  id INTEGER PRIMARY KEY,
  'user_id' INTEGER NOT NULL,
  'reviewer_id' INTEGER NOT NULL,
  'name' TEXT NOT NULL,
  'type' TEXT NOT NULL,
  'accepted' INTEGER,
  'address' TEXT,
  'reviewer_comment' TEXT,
  'created_at' DATE NOT NULL,
  'updated_at' DATE NOT NULL,

  FOREIGN KEY('user_id') REFERENCES 'users'('id'),
  FOREIGN KEY('reviewer_id') REFERENCES 'users'('id')
);
