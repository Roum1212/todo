package get_reminderByID_quary

import (
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

type Command struct {
	reminderID reminder_id_model.ReminderID
}

func NewCommand(reminderID reminder_id_model.ReminderID) Command {
	return Command{
		reminderID: reminderID,
	}
}
