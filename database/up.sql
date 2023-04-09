-- users

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  auth_provider VARCHAR(255) NOT NULL,
  hash_pass VARCHAR(255) NOT NULL,
  auth_token VARCHAR(255)
);
