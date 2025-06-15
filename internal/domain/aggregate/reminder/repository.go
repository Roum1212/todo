package reminder_aggregate

import (
	"context"
	"errors"

	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

var ErrReminderNotFound = errors.New("reminder not found")

//go:generate minimock -i ReminderRepository -g -o ./mock -p reminder_aggregate_mock -s "_minimock.go"
type ReminderRepository interface {
	DeleteReminder(ctx context.Context, reminderID reminder_id_model.ReminderID) error
	GetAllReminders(ctx context.Context) ([]Reminder, error)
	GetReminderByID(ctx context.Context, reminderID reminder_id_model.ReminderID) (Reminder, error)
	SaveReminder(ctx context.Context, reminder Reminder) error
}
