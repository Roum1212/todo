package reminder_aggregate

import (
	"context"
)

type ReminderRepository interface {
	SaveReminder(ctx context.Context, reminder Reminder) error
	DeleteReminder(ctx context.Context, reminderID Reminder) error
}
