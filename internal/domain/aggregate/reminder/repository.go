package reminder_aggregate

//go:generate minimock -i ReminderRepository -o mock/ -s "_mock.go"

import (
	"context"
	"errors"

	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

var ErrRemindersNotFound = errors.New("reminders not found")

type ReminderRepository interface {
	SaveReminder(ctx context.Context, reminder Reminder) error
	DeleteReminder(ctx context.Context, reminderID reminder_id_model.ReminderID) error
	GetReminderByID(ctx context.Context, reminderID reminder_id_model.ReminderID) (Reminder, error)
	GetAllReminders(ctx context.Context) ([]Reminder, error)
}
