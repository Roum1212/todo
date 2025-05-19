package get_all_reminders_http_handler

import (
	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
)

type Reminder struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewReminderDTOs(reminderSlice []reminder_aggregate.Reminder) []Reminder {
	var reminderDTOs []Reminder

	for i := range reminderSlice {
		reminderDTOs = append(reminderDTOs, Reminder{
			ID:          int(reminderSlice[i].GetID()),
			Title:       string(reminderSlice[i].GetTitle()),
			Description: string(reminderSlice[i].GetDescription()),
		})
	}

	return reminderDTOs
}
