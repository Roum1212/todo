package reminder_aggregate

import (
	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
)

type Reminder struct {
	Id          reminder_id_model.ReminderID
	Title       reminder_title_model.ReminderTitle
	Description reminder_description_model.ReminderDescription
}

func (x Reminder) GetID() reminder_id_model.ReminderID {
	return x.Id
}

func (x Reminder) GetTitle() reminder_title_model.ReminderTitle {
	return x.Title
}

func (x Reminder) GetDescription() reminder_description_model.ReminderDescription {
	return x.Description
}

func NewReminder(
	id reminder_id_model.ReminderID,
	title reminder_title_model.ReminderTitle,
	description reminder_description_model.ReminderDescription,
) Reminder {
	return Reminder{
		Id:          id,
		Title:       title,
		Description: description,
	}
}
