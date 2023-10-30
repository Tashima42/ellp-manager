CREATE TABLE IF NOT EXISTS documents (
  id INTEGER PRIMARY KEY,
  'user_id' INTEGER NOT NULL,
  'requester_id' INTEGER NOT NULL,
  'name' TEXT NOT NULL,
  'type' TEXT NOT NULL,
  'accepted' INTEGER,
  'address' TEXT,
  'requester_comment' TEXT,
  'created_at' DATE NOT NULL,
  'updated_at' DATE NOT NULL,

  FOREIGN KEY('user_id') REFERENCES 'users'('id'),
  FOREIGN KEY('requester_id') REFERENCES 'users'('id')
);
