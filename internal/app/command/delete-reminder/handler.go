package delete_reminder_command

import (
	"context"
	"fmt"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
)

type Handler struct {
	repository reminder_aggregate.RepositoryDelete
}

func (x Handler) Handle(ctx context.Context, d DeleteCommand) error {

	reminder := reminder_aggregate.DeleteReminder(d.id)

	if err := x.repository.DeleteReminder(ctx, reminder); err != nil {
		return fmt.Errorf("failed to delete reminder: %w", err)
	}

	return nil
}

func DeleteHandler(repository reminder_aggregate.RepositoryDelete) Handler {
	return Handler{
		repository: repository,
	}
}
