package create_reminder_command

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

const tracerName = "create_reminder_command"

//go:generate minimock -i CommandHandler -g -o ./mock -p create_reminder_command_mock -s "_minimock.go"
type CommandHandler interface {
	HandleCommand(ctx context.Context, command Command) error
}

type commandHandler struct {
	repository reminder_aggregate.ReminderRepository
}

func (x commandHandler) HandleCommand(ctx context.Context, c Command) error {
	tracer := otel.Tracer(tracerName)
	_, span := tracer.Start(ctx, "Repository.SaveReminder")

	defer span.End()

	reminderID := reminder_id_model.GenerateReminderID()
	reminder := reminder_aggregate.NewReminder(reminderID, c.title, c.description)

	if err := x.repository.SaveReminder(ctx, reminder); err != nil {
		return fmt.Errorf("failed to save reminder: %w", err)
	}

	return nil
}

func NewCommandHandler(repository reminder_aggregate.ReminderRepository) CommandHandler {
	return commandHandler{
		repository: repository,
	}
}
