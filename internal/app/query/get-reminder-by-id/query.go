package get_reminder_by_id_query

import (
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

type Query struct {
	reminderID reminder_id_model.ReminderID
}

func NewQuery(reminderID reminder_id_model.ReminderID) Query {
	return Query{
		reminderID: reminderID,
	}
}
