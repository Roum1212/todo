-- +goose Up
CREATE TABLE IF NOT EXISTS reminders
(
    id          int PRIMARY KEY,
    title       text NOT NULL,
    description text NOT NULL
);

CREATE TABLE IF NOT EXISTS accounts
(
    login text NOT NULL,
    password text NOT NULL
);


-- +goose Down
DROP TABLE reminders;
