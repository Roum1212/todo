package reminder_aggregate

import (
	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
)

type Reminder struct {
	id          reminder_id_model.ReminderID
	title       reminder_title_model.ReminderTitle
	description reminder_description_model.ReminderDescription
	delete      bool
}

func (x Reminder) GetID() reminder_id_model.ReminderID {
	return x.id
}

func (x Reminder) GetTitle() reminder_title_model.ReminderTitle {
	return x.title
}

func (x Reminder) GetDescription() reminder_description_model.ReminderDescription {
	return x.description
}

func NewReminder(
	id reminder_id_model.ReminderID,
	title reminder_title_model.ReminderTitle,
	description reminder_description_model.ReminderDescription,
) Reminder {
	return Reminder{
		id:          id,
		title:       title,
		description: description,
		delete:      false,
	}
}

func DeleteReminder(id reminder_id_model.ReminderID) Reminder {
	return Reminder{
		id:          id,
		title:       "-",
		description: "-",
		delete:      true,
	}
}
