package create_reminder_command

//go:generate minimock -i CommandHandler -o mock/ -s "_mock.go"

import (
	"context"
	"fmt"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

type CommandHandler interface {
	HandleCommand(ctx context.Context, command Command) error
}

type commandHandler struct {
	repository reminder_aggregate.ReminderRepository
}

func (x commandHandler) HandleCommand(ctx context.Context, c Command) error {
	reminderID := reminder_id_model.GenerateReminderID()
	reminder := reminder_aggregate.NewReminder(reminderID, c.title, c.description)

	if err := x.repository.SaveReminder(ctx, reminder); err != nil {
		return fmt.Errorf("failed to save reminder: %w", err)
	}

	return nil
}

func NewHandler(repository reminder_aggregate.ReminderRepository) CommandHandler {
	return commandHandler{
		repository: repository,
	}
}
