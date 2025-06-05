package get_reminder_by_id_query

import (
	"context"
	"errors"
	"fmt"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
)

var ErrReminderNotFound = errors.New("reminder not found")

//go:generate minimock -i QueryHandler -g -o ./mock -p get_reminder_by_id_query_mock -s "_minimock.go"
type QueryHandler interface {
	HandleQuery(ctx context.Context, q Query) (reminder_aggregate.Reminder, error)
}

type queryHandler struct {
	repository reminder_aggregate.ReminderRepository
}

func (x queryHandler) HandleQuery(ctx context.Context, q Query) (reminder_aggregate.Reminder, error) {
	reminder, err := x.repository.GetReminderByID(ctx, q.id)
	if err != nil {
		switch {
		case errors.Is(err, reminder_aggregate.ErrReminderNotFound):
			return reminder_aggregate.Reminder{}, ErrReminderNotFound
		default:
			return reminder_aggregate.Reminder{}, fmt.Errorf("failed to get reminder: %w", err)
		}
	}

	return reminder, nil
}

func NewQueryHandler(repository reminder_aggregate.ReminderRepository) QueryHandler {
	return queryHandler{
		repository: repository,
	}
}
