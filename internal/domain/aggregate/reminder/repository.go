package reminder_aggregate

import (
	"context"
)

type RepositorySave interface {
	SaveReminder(ctx context.Context, reminder Reminder) error
}

type RepositoryDelete interface {
	DeleteReminder(ctx context.Context, reminder Reminder) error
}
