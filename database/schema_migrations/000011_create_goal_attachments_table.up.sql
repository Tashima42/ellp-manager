CREATE TABLE IF NOT EXISTS goal_attachments (
  id INTEGER PRIMARY KEY,
  'user_id' INTEGER NOT NULL,
  'goal_id' INTEGER NOT NULL,
  'name' TEXT NOT NULL,
  'comment' TEXT,
  'document_id' INTEGER,
  'created_at' DATE NOT NULL,
  'updated_at' DATE NOT NULL,

  FOREIGN KEY('user_id') REFERENCES 'users'('id'),
  FOREIGN KEY('goal_id') REFERENCES 'goals'('id'),
  FOREIGN KEY('document_id') REFERENCES 'documents'('id')
);
