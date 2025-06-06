package get_reminder_by_id_http_handler

import (
	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
)

type Reminder struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewReminder(reminder reminder_aggregate.Reminder) Reminder {
	return Reminder{
		ID:          int64(reminder.GetID()),
		Title:       string(reminder.GetTitle()),
		Description: string(reminder.GetDescription()),
	}
}
