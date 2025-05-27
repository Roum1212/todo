-- +goose Up
CREATE TABLE accounts
(
    login    TEXT PRIMARY KEY,
    password TEXT NOT NULL
);


-- +goose Down
DROP TABLE accounts;
