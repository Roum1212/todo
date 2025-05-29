package postgresql_reminder_repository

import (
	"fmt"

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

func NewReminders(reminderDTOs []Reminder) ([]reminder_aggregate.Reminder, error) {
	reminders := make([]reminder_aggregate.Reminder, len(reminderDTOs))

	for i := range reminderDTOs {
		reminderID := reminder_id_model.ReminderID(reminderDTOs[i].ID)

		reminderTitle, err := reminder_title_model.NewReminderTitle(reminderDTOs[i].Title)
		if err != nil {
			return nil, fmt.Errorf("faild to create reminder: %w", err)
		}

		reminderDescription, err := reminder_description_model.NewReminderDescription(reminderDTOs[i].Description)
		if err != nil {
			return nil, fmt.Errorf("faild to create reminder: %w", err)
		}

		reminders[i] = reminder_aggregate.NewReminder(reminderID, reminderTitle, reminderDescription)
	}

	return reminders, nil
}
