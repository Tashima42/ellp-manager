CREATE TABLE IF NOT EXISTS workshop_user (
  id INTEGER PRIMARY KEY,
  'workshop_id' INTEGER NOT NULL,
  'user_id' INTEGER NOT NULL,
  'role' INTEGER NOT NULL,
  'created_at' DATE NOT NULL,
  'updated_at' DATE NOT NULL,

  FOREIGN KEY('workshop_id') REFERENCES 'workshops'('id')
  FOREIGN KEY('user_id') REFERENCES 'users'('id')
);
