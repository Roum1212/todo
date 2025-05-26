-- +goose Up
CREATE TABLE reminders
(
    id          int PRIMARY KEY,
    title       text NOT NULL,
    description text NOT NULL
);

-- +goose Down
DROP TABLE reminders;
