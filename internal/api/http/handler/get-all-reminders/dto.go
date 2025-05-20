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
	reminderDTOs := make([]Reminder, len(reminders))
	for i := range reminders {
		reminderDTOs[i] = Reminder{
			ID:          int(reminders[i].GetID()),
			Title:       string(reminders[i].GetTitle()),
			Description: string(reminders[i].GetDescription()),
		}
	}

	return reminderDTOs
}
