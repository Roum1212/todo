package get_all_reminders_quer

import (
	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
)

type Query struct {
	title       reminder_title_model.ReminderTitle
	description reminder_description_model.ReminderDescription
	id          reminder_id_model.ReminderID
}

func NewQuery(q Query) Query {
	return Query{
		title:       q.title,
		description: q.description,
		id:          q.id,
	}
}
