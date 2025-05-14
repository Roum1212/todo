package reminder_aggregate

import (
	"context"

	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

type ReminderRepository interface {
	SaveReminder(ctx context.Context, reminder Reminder) error
	DeleteReminder(ctx context.Context, reminderID reminder_id_model.ReminderID) error
	GetReminderByID(ctx context.Context, reminderID reminder_id_model.ReminderID) (Reminder, error)
}
