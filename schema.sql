CREATE TABLE IF NOT EXISTS endpoint (
  id INTEGER PRIMARY KEY AUTOINCREMENT,

  method VARCHAR(10) NOT NULL,
  url VARCHAR(255) NOT NULL,
  query_params TEXT,
  headers TEXT,
  request_body_type VARCHAR(50)
  request_body TEXT,

  response_body_type VARCHAR(50)
  request_body TEXT,
  status_code INTEGER,
)
