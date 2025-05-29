package get_all_reminders_quer

import (
	"context"
	"errors"
	"fmt"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
)

var ErrRemindersNotFound = errors.New("reminders not found")

type Handler struct {
	repository reminder_aggregate.ReminderRepository
}

func (x Handler) Handle(ctx context.Context) ([]reminder_aggregate.Reminder, error) {
	reminders, err := x.repository.GetAllReminders(ctx)
	if err != nil {
		switch {
		case errors.Is(err, reminder_aggregate.ErrRemindersNotFound):
			return nil, ErrRemindersNotFound
		default:
			return nil, fmt.Errorf("failed to get all reminders: %w", err)
		}
	}

	return reminders, nil
}

func NewHandler(repository reminder_aggregate.ReminderRepository) Handler {
	return Handler{
		repository: repository,
	}
}
