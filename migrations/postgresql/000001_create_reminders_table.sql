-- +goose Up
CREATE TABLE IF NOT EXISTS reminders
(
    id          int PRIMARY KEY,
    title       text NOT NULL,
    description text NOT NULL
);


-- +goose Down
DROP TABLE reminders;
