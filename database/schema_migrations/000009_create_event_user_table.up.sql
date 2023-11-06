CREATE TABLE IF NOT EXISTS event_user (
  id INTEGER PRIMARY KEY,
  'event_id' INTEGER NOT NULL,
  'user_id' INTEGER NOT NULL,
  'participated' INTEGER,
  'created_at' DATE NOT NULL,
  'updated_at' DATE NOT NULL,

  FOREIGN KEY('event_id') REFERENCES 'events'('id'),
  FOREIGN KEY('user_id') REFERENCES 'users'('id')
);
