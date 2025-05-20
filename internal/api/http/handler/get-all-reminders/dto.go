package get_all_reminders_http_handler

import (
	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
)

type Reminder struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewReminderDTOs(remindersSlice []reminder_aggregate.Reminder) []Reminder {
	var reminderDTOs []Reminder

	for i := range remindersSlice {
		reminderDTOs = append(reminderDTOs, Reminder{
			ID:          int(remindersSlice[i].GetID()),
			Title:       string(remindersSlice[i].GetTitle()),
			Description: string(remindersSlice[i].GetDescription()),
		})
	}

	return reminderDTOs
}
