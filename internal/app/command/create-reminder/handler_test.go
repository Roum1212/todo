package create_reminder_command

import (
	"context"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
	"github.com/Roum1212/todo/internal/domain/aggregate/reminder/mock"
)

func TestCommandHandler_HandleCommand_Success(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	command := Command{
		title:       "title",
		description: "description",
	}

	reminderRepositoryMock := mock.NewReminderRepositoryMock(mc).
		SaveReminderMock.
		Inspect(func(ctx context.Context, reminder reminder_aggregate.Reminder) {
			require.Equal(t, command.title, reminder.GetTitle())
			require.Equal(t, command.description, reminder.GetDescription())
		}).
		Return(nil)

	handler := NewHandler(reminderRepositoryMock)

	err := handler.HandleCommand(context.Background(), command)
	require.NoError(t, err)
}

func TestCommandHandler_HandleCommand_Error(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	command := Command{
		title:       "title",
		description: "description",
	}

	reminderRepositoryMock := mock.NewReminderRepositoryMock(mc).
		SaveReminderMock.
		Inspect(func(ctx context.Context, reminder reminder_aggregate.Reminder) {
			require.Equal(t, command.title, reminder.GetTitle())
		}).
		Return(assert.AnError)

	handler := NewHandler(reminderRepositoryMock)

	err := handler.HandleCommand(context.Background(), command)
	require.ErrorIs(t, err, assert.AnError)
}
