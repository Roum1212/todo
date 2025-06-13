package create_reminder_command

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
)

const tracerName = "github.com/Roum1212/todo/internal/app/command/create-reminder"

//go:generate minimock -i CommandHandler -g -o ./mock -p create_reminder_command_mock -s "_minimock.go"
type CommandHandler interface {
	HandleCommand(ctx context.Context, command Command) error
}

type commandHandler struct {
	repository reminder_aggregate.ReminderRepository
}

func (x commandHandler) HandleCommand(ctx context.Context, c Command) error {
	if err := x.repository.SaveReminder(
		ctx,
		reminder_aggregate.NewReminder(c.id, c.title, c.description),
	); err != nil {
		return fmt.Errorf("failed to save reminder: %w", err)
	}

	return nil
}

func NewCommandHandler(repository reminder_aggregate.ReminderRepository) CommandHandler {
	return commandHandler{
		repository: repository,
	}
}

type tracerCommandHandler struct {
	commandHandler CommandHandler
	tracer         trace.Tracer
}

func (x tracerCommandHandler) HandleCommand(ctx context.Context, c Command) error {
	_, span := x.tracer.Start(ctx, "commandHandler.HandleCommand")
	defer span.End()

	if err := x.commandHandler.HandleCommand(ctx, c); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return err
	}

	return nil
}

func NewCommandHandlerWithTracing(commandHandler CommandHandler) CommandHandler {
	return tracerCommandHandler{
		commandHandler: commandHandler,
		tracer:         otel.Tracer(tracerName),
	}
}
