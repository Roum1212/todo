package postgresql_reminder_repository

import (
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

func NewReminderID(reminderID reminder_id_model.ReminderID) string {
	return string(reminderID)
}
