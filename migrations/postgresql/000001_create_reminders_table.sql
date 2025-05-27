-- +goose Up
CREATE TABLE reminders
(
    id          INT PRIMARY KEY,
    title       TEXT NOT NULL,
    description TEXT NOT NULL
);

-- +goose Down
DROP TABLE reminders;
