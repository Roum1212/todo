package get_all_reminders_query

import (
	"context"
	"errors"
	"fmt"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
)

var ErrRemindersNotFound = errors.New("reminders not found")

//go:generate minimock -i QueryHandler -o mock/ -s "_mock.go"
type QueryHandler interface {
	HandleQuery(ctx context.Context) ([]reminder_aggregate.Reminder, error)
}

type queryHandler struct {
	repository reminder_aggregate.ReminderRepository
}

func (x queryHandler) HandleQuery(ctx context.Context) ([]reminder_aggregate.Reminder, error) {
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

func NewHandler(repository reminder_aggregate.ReminderRepository) QueryHandler {
	return queryHandler{
		repository: repository,
	}
}
