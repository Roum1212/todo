-- +goose Up
CREATE TABLE accounts
(
    login    text PRIMARY KEY NOT NULL,
    password text             NOT NULL
);


-- +goose Down
DROP TABLE accounts;
