-- +goose Up
CREATE TABLE feeds (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name VARCHAR(50) UNIQUE NOT NULL,
  url VARCHAR UNIQUE, 
  user_id UUID,
  FOREIGN KEY (user_id) REFERENCES users(id)
    ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds
