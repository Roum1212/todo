package postgresql_reminder_repository

import (
	"context"
	"log"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

type Repository struct{}

func (x Repository) SaveReminder(
	ctx context.Context,
	reminder reminder_aggregate.Reminder,
) error {
	log.Println("id", reminder.GetID())
	log.Println("title", reminder.GetTitle())
	log.Println("description", reminder.GetDescription())

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

func NewRepository() Repository {
	return Repository{}
}
