package delete_reminder_command

import (
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

type DeleteCommand struct {
	reminderID reminder_id_model.ReminderID
}

func NewDeleteCommand(id reminder_id_model.ReminderID) DeleteCommand {
	return DeleteCommand{
		reminderID: id,
	}
}
