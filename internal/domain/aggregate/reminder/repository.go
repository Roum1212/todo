package reminder_aggregate

import (
	"context"
)

type Repository interface {
	SaveReminder(ctx context.Context, reminder Reminder) error
}
