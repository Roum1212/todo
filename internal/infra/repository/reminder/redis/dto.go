package redis_reminder_repository

import (
	"fmt"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
	reminder_v1 "github.com/Roum1212/todo/pkg/gen/reminder/v1"
)

func ToReminder(reminderDTO *reminder_v1.Reminder) (reminder_aggregate.Reminder, error) {
	id, err := reminder_id_model.NewReminderID(reminderDTO.GetId())
	if err != nil {
		return reminder_aggregate.Reminder{}, fmt.Errorf("failed to create id: %w", err)
	}

	title, err := reminder_title_model.NewReminderTitle(reminderDTO.GetTitle())
	if err != nil {
		return reminder_aggregate.Reminder{}, fmt.Errorf("failed to create title: %w", err)
	}

	description, err := reminder_description_model.NewReminderDescription(reminderDTO.GetDescription())
	if err != nil {
		return reminder_aggregate.Reminder{}, fmt.Errorf("failed to create description: %w", err)
	}

	return reminder_aggregate.NewReminder(id, title, description), nil
}

func ToReminders(reminderDTOs ...*reminder_v1.Reminder) ([]reminder_aggregate.Reminder, error) {
	reminders := make([]reminder_aggregate.Reminder, len(reminderDTOs))

	for i := range reminderDTOs {
		reminder, err := ToReminder(reminderDTOs[i])
		if err != nil {
			return nil, fmt.Errorf("failed to create reminder: %w", err)
		}

		reminders[i] = reminder
	}

	return reminders, nil
}

func NewReminderDTO(reminder reminder_aggregate.Reminder) *reminder_v1.Reminder {
	return &reminder_v1.Reminder{
		Id:          int64(reminder.GetID()),
		Title:       string(reminder.GetTitle()),
		Description: string(reminder.GetDescription()),
	}
}
