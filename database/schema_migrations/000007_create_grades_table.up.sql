CREATE TABLE IF NOT EXISTS grades (
  id INTEGER PRIMARY KEY,
  'workshop_id' INTEGER NOT NULL,
  'user_id' INTEGER NOT NULL,
  'grade' INTEGER NOT NULL,
  'created_at' DATE NOT NULL,
  'updated_at' DATE NOT NULL,

  FOREIGN KEY('workshop_id') REFERENCES 'workshops'('id')
  FOREIGN KEY('user_id') REFERENCES 'users'('id')
);
