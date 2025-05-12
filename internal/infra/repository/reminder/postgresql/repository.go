package postgresql_reminder_repository

import (
	"context"
	"log"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
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
	reminder reminder_aggregate.Reminder,
) error {
	log.Println("The reminder with the ID", reminder.GetID(), "has been deleted")

	return nil
}

func NewRepository() Repository {
	return Repository{}
}
