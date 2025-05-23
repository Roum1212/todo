package delete_reminder_command

import (
	"context"
	"fmt"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
)

type Handler struct {
	repository reminder_aggregate.ReminderRepository
}

func (x Handler) Handle(ctx context.Context, c Command) error {
	if err := x.repository.DeleteReminder(ctx, c.reminderID); err != nil {
		return fmt.Errorf("failed to delete reminder: %w", err)
	}

	return nil
}

func NewHandler(repository reminder_aggregate.ReminderRepository) Handler {
	return Handler{
		repository: repository,
	}
}
