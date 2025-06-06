package delete_reminder_command

import (
	"fmt"

	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

type Command struct {
	id reminder_id_model.ReminderID
}

func (x Command) GetID() reminder_id_model.ReminderID {
	return x.id
}

func (x Command) Validate() error {
	if err := x.id.Validate(); err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	return nil
}

func NewCommand(reminderID reminder_id_model.ReminderID) Command {
	return Command{
		id: reminderID,
	}
}
