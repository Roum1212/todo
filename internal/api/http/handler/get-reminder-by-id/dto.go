package get_reminder_by_id_http_handler

import (
	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
)

type Reminder struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewReminder(
	id reminder_id_model.ReminderID,
	title reminder_title_model.ReminderTitle,
	description reminder_description_model.ReminderDescription,
) Reminder {
	return Reminder{
		ID:          int(id),
		Title:       string(title),
		Description: string(description),
	}
}
