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
	reminder, err := x.repository.GetAllReminders(ctx)
	if err != nil {
		return []reminder_aggregate.Reminder{}, fmt.Errorf("reminder_id cannot be parsed as int: %w", err)
	}

	return reminder, nil
}

func NewHandler(repository reminder_aggregate.ReminderRepository) Handler {
	return Handler{
		repository: repository}
}
