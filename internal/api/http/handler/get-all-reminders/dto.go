package get_all_reminders_http_handler

import (
	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
)

type Reminder struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
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
