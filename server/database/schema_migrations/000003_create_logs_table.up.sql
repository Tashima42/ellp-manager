CREATE TABLE IF NOT EXISTS logs (
  id INTEGER PRIMARY KEY,
  'user_id' INTEGER NOT NULL,
  'document_id' INTEGER NOT NULL,
  'action' TEXT NOT NULL,
  'description' TEXT,
  'created_at' DATE NOT NULL,
  'updated_at' DATE NOT NULL,

  FOREIGN KEY('user_id') REFERENCES 'users'('id'),
  FOREIGN KEY('document_id') REFERENCES 'documents'('id')
);
