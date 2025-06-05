package reminder_aggregate

import (
	"context"
	"errors"

	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

var ErrReminderNotFound = errors.New("reminder not found")

//go:generate minimock -i ReminderRepository -o mock/ -s "_mock.go"
type ReminderRepository interface {
	DeleteReminder(ctx context.Context, reminderID reminder_id_model.ReminderID) error
	GetReminderByID(ctx context.Context, reminderID reminder_id_model.ReminderID) (Reminder, error)
	GetAllReminders(ctx context.Context) ([]Reminder, error)
	SaveReminder(ctx context.Context, reminder Reminder) error
}
