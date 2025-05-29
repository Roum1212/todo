package get_reminder_by_id_query

import (
	"context"
	"errors"
	"fmt"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
)

var ErrReminderNotFound = errors.New("reminder not found")

type Handler struct {
	repository reminder_aggregate.ReminderRepository
}

func (x Handler) Handle(ctx context.Context, q Query) (reminder_aggregate.Reminder, error) {
	reminder, err := x.repository.GetReminderByID(ctx, q.reminderID)
	if err != nil {
		switch {
		case errors.Is(err, ErrReminderNotFound):
			return reminder_aggregate.Reminder{}, ErrReminderNotFound
		default:
			return reminder_aggregate.Reminder{}, fmt.Errorf("failed to get reminder: %w", err)
		}
	}

	return reminder, nil
}

func NewHandler(repository reminder_aggregate.ReminderRepository) Handler {
	return Handler{
		repository: repository,
	}
}
