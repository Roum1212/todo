package get_all_reminders_query

import (
	"context"
	"errors"
	"fmt"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
)

var ErrReminderNotFound = errors.New("reminders not found")

//go:generate minimock -i QueryHandler -g -o ./mock -p get_all_reminders_query_mock -s "_minimock.go"
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
		case errors.Is(err, reminder_aggregate.ErrReminderNotFound):
			return nil, ErrReminderNotFound
		default:
			return nil, fmt.Errorf("failed to get all reminders: %w", err)
		}
	}

	return reminders, nil
}

func NewQueryHandler(repository reminder_aggregate.ReminderRepository) QueryHandler {
	return queryHandler{
		repository: repository,
	}
}
