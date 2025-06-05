package delete_reminder_command

import (
	"fmt"

	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

type Command struct {
	reminderID reminder_id_model.ReminderID
}

func (x Command) GetReminderID() reminder_id_model.ReminderID {
	return x.reminderID
}

func (x Command) Validate() error {
	if err := x.reminderID.Validate(); err != nil {
		return fmt.Errorf("invalid reminder id: %w", err)
	}

	return nil
}

func NewCommand(reminderID reminder_id_model.ReminderID) Command {
	return Command{
		reminderID: reminderID,
	}
}
