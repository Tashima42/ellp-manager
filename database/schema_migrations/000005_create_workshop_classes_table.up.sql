CREATE TABLE IF NOT EXISTS workshop_classes (
  id INTEGER PRIMARY KEY,
  'workshop_id' INTEGER NOT NULL,
  'name' TEXT NOT NULL,
  'description' TEXT,
  'number' iNTEGER NOT NULL,
  'start_at' DATE NOT NULL,
  'end_at' DATE NOT NULL,
  'created_at' DATE NOT NULL,
  'updated_at' DATE NOT NULL,

  FOREIGN KEY('workshop_id') REFERENCES 'workshops'('id')
);
