package get_reminder_by_id_query

import (
	"fmt"

	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

type Query struct {
	id reminder_id_model.ReminderID
}

func (x Query) GetID() reminder_id_model.ReminderID {
	return x.id
}

func (x Query) Validate() error {
	if err := x.id.Validate(); err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	return nil
}

func NewQuery(reminderID reminder_id_model.ReminderID) Query {
	return Query{
		id: reminderID,
	}
}
