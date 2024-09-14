-- +goose Up
CREATE TABLE chat (
      id SERIAL PRIMARY KEY,
      title TEXT
);

-- +goose Down
DROP TABLE chat;
