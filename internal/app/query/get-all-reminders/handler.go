package get_all_reminders_quer

import (
	"context"
	"fmt"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
)

type Handler struct {
	repository reminder_aggregate.ReminderRepository
}

func (x Handler) Handle(ctx context.Context) ([]reminder_aggregate.Reminder, error) {
	reminders, err := x.repository.GetAllReminders(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all reminders: %w", err)
	}

	return reminders, nil
}

func NewHandler(repository reminder_aggregate.ReminderRepository) Handler {
	return Handler{
		repository: repository,
	}
}
