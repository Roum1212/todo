package get_reminderByID_quary

import (
	"context"
	"fmt"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
)

type Handler struct {
	repository reminder_aggregate.ReminderRepository
}

func (x Handler) Handle(ctx context.Context, c Command) (reminder_aggregate.Reminder, error) {
	reminderID, err := x.repository.GetReminderByID(ctx, c.reminderID)
	if err != nil {
		return reminder_aggregate.Reminder{}, fmt.Errorf("failed to get reminder: %w", err)
	}

	return reminderID, nil
}

func NewHandler(repository reminder_aggregate.ReminderRepository) Handler {
	return Handler{
		repository: repository,
	}
}
