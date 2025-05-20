package get_all_reminders_http_handler

import (
	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
)

type Reminder struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewReminderDTOs(reminders []reminder_aggregate.Reminder) []Reminder {
	var reminderDTOs []Reminder

	for i := range reminders {
		reminderDTOs = append(reminderDTOs, Reminder{
			ID:          int(reminders[i].GetID()),
			Title:       string(reminders[i].GetTitle()),
			Description: string(reminders[i].GetDescription()),
		})
	}

	return reminderDTOs
}
