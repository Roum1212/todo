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

func NewRemindersSlice(remindersDTOs []Reminder, remindersSlice []reminder_aggregate.Reminder) []reminder_aggregate.Reminder {
	for i := range remindersDTOs {
		remindersSlice[i] = reminder_aggregate.NewReminder(
			reminder_id_model.ReminderID(remindersDTOs[i].ID),
			reminder_title_model.NewReminderTitle(remindersDTOs[i].Title),
			reminder_description_model.NewReminderDescription(remindersDTOs[i].Description),
		)
	}

	return remindersSlice
}
