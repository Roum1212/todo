package postgresql_reminder_repository

import (
	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
)

type Reminder struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewReminderJson(
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

func NewReminderJsonSlice(slice []reminder_aggregate.Reminder) []Reminder {
	var reminders []Reminder

	for _, value := range slice {
		reminders = append(reminders, Reminder{
			ID:          int(value.GetID()),
			Title:       string(value.GetTitle()),
			Description: string(value.GetDescription()),
		})
	}

	return reminders
}
