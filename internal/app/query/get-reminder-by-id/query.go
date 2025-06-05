package get_reminder_by_id_query

import (
	"fmt"

	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

type Query struct {
	reminderID reminder_id_model.ReminderID
}

func (x Query) GetReminderID() reminder_id_model.ReminderID {
	return x.reminderID
}

func (x Query) Validate() error {
	if err := x.reminderID.Validate(); err != nil {
		return fmt.Errorf("invalid reminder id: %w", err)
	}

	return nil
}

func NewQuery(reminderID reminder_id_model.ReminderID) Query {
	return Query{
		reminderID: reminderID,
	}
}
