package postgresql_reminder_repository

import (
	"fmt"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
	reminder_v1 "github.com/Roum1212/todo/pkg/gen/reminder/v1"
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

func NewReminders(reminders ...reminder_aggregate.Reminder) []Reminder {
	reminderDTOs := make([]Reminder, len(reminders))

	for i := range reminders {
		reminderDTOs[i] = NewReminder(reminders[i])
	}

	return reminderDTOs
}

func ToReminder(reminderDTO Reminder) (reminder_aggregate.Reminder, error) {
	id, err := reminder_id_model.NewReminderID(reminderDTO.ID)
	if err != nil {
		return reminder_aggregate.Reminder{}, fmt.Errorf("failed to create reminder id: %w", err)
	}

	title, err := reminder_title_model.NewReminderTitle(reminderDTO.Title)
	if err != nil {
		return reminder_aggregate.Reminder{}, fmt.Errorf("failed to create reminder title: %w", err)
	}

	description, err := reminder_description_model.NewReminderDescription(reminderDTO.Description)
	if err != nil {
		return reminder_aggregate.Reminder{}, fmt.Errorf("failed to create reminder description: %w", err)
	}

	return reminder_aggregate.NewReminder(id, title, description), nil
}

func ToReminders(reminderDTOs ...Reminder) ([]reminder_aggregate.Reminder, error) {
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

func ToReminderFromProto(reminderRPC *reminder_v1.Reminder) (reminder_aggregate.Reminder, error) {
	id, err := reminder_id_model.NewReminderID(reminderRPC.GetId())
	if err != nil {
		return reminder_aggregate.Reminder{}, fmt.Errorf("failed to create id: %w", err)
	}

	title, err := reminder_title_model.NewReminderTitle(reminderRPC.GetTitle())
	if err != nil {
		return reminder_aggregate.Reminder{}, fmt.Errorf("failed to create title: %w", err)
	}

	description, err := reminder_description_model.NewReminderDescription(reminderRPC.GetDescription())
	if err != nil {
		return reminder_aggregate.Reminder{}, fmt.Errorf("failed to create description: %w", err)
	}

	reminder := reminder_aggregate.NewReminder(id, title, description)

	return reminder, nil
}

func NewProtoReminder(reminder reminder_aggregate.Reminder) *reminder_v1.Reminder {
	return &reminder_v1.Reminder{
		Id:          int64(reminder.GetID()),
		Title:       string(reminder.GetTitle()),
		Description: string(reminder.GetDescription()),
	}
}
