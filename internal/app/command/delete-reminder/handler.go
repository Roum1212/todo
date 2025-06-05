package delete_reminder_command

import (
	"context"
	"fmt"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
)

//go:generate minimock -i CommandHandler -g -o ./mock -p delete_reminder_command_mock -s "_minimock.go"
type CommandHandler interface {
	HandleCommand(ctx context.Context, c Command) error
}

type commandHandler struct {
	repository reminder_aggregate.ReminderRepository
}

func (x commandHandler) HandleCommand(ctx context.Context, c Command) error {
	if err := x.repository.DeleteReminder(ctx, c.id); err != nil {
		return fmt.Errorf("failed to delete reminder: %w", err)
	}

	return nil
}

func NewCommandHandler(repository reminder_aggregate.ReminderRepository) CommandHandler {
	return commandHandler{
		repository: repository,
	}
}
