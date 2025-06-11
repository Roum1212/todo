package get_all_reminders_query

import (
	"context"
	"errors"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
)

var ErrReminderNotFound = errors.New("reminders not found")

const tracerName = "github.com/Roum1212/todo/internal/app/query/get-all-reminders"

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

type tracerQueryHandler struct {
	queryHandler QueryHandler
	tracer       trace.Tracer
}

func (x tracerQueryHandler) HandleQuery(ctx context.Context) ([]reminder_aggregate.Reminder, error) {
	_, span := x.tracer.Start(ctx, "QueryHandler.HandleQuery")
	defer span.End()

	reminders, err := x.queryHandler.HandleQuery(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return nil, err
	}

	return reminders, nil
}

func NewQueryHandlerTracer(queryHandler QueryHandler) QueryHandler {
	return tracerQueryHandler{
		queryHandler: queryHandler,
		tracer:       otel.Tracer(tracerName),
	}
}
