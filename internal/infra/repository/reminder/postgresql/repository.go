package postgresql_reminder_repository

import (
	"context"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

type Repository struct {
	db *pgxpool.Pool
}

func (x Repository) SaveReminder(
	ctx context.Context,
	reminder reminder_aggregate.Reminder,
) error {
	sql, args, err := squirrel.
		Insert("reminders").
		Columns("reminder_id", "reminder_title", "reminder_description").
		Values(int(reminder.GetID()),
			string(reminder.GetTitle()),
			string(reminder.GetDescription())).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = x.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	log.Println("reminder saved")

	return nil
}

func (x Repository) DeleteReminder(
	ctx context.Context,
	reminderID reminder_id_model.ReminderID,
) error {
	log.Println("The reminder with the ID", reminderID, "has been deleted")

	return nil
}

func (x Repository) GetReminderByID(
	ctx context.Context,
	reminderID reminder_id_model.ReminderID,
) (reminder_aggregate.Reminder, error) {
	log.Println("The reminder with the ID", reminderID, "has been found")

	reminder := reminder_aggregate.Reminder{}
	return reminder, nil
}

func (x Repository) GetAllReminders(ctx context.Context) ([]reminder_aggregate.Reminder, error) {
	log.Println("Getting all reminders", ctx)

	return []reminder_aggregate.Reminder{}, nil
}

func NewRepository(db *pgxpool.Pool) Repository {
	return Repository{
		db: db,
	}
}
