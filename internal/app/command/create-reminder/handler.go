package create_reminder_command

import (
	"context"
	"fmt"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

type Handler struct {
	repository reminder_aggregate.ReminderRepository
}

func (x Handler) Handle(ctx context.Context, c CreateCommand) error {
	reminderID := reminder_id_model.GenerateReminderID()

	reminder := reminder_aggregate.NewReminder(reminderID, c.title, c.description)

	if err := x.repository.SaveReminder(ctx, reminder); err != nil {
		return fmt.Errorf("failed to save reminder: %w", err)
	}

	return nil
}

func NewHandler(repository reminder_aggregate.ReminderRepository) Handler {
	return Handler{
		repository: repository,
	}
}
