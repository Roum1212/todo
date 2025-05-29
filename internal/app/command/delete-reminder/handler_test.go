package delete_reminder_command

import (
	"context"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Roum1212/todo/internal/domain/aggregate/reminder/mock"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

func TestCommandHandler_HandleCommand_Success(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	command := Command{
		reminderID: 123,
	}

	reminderRepositoryMock := mock.NewReminderRepositoryMock(mc).
		DeleteReminderMock.
		Inspect(func(ctx context.Context, reminderID reminder_id_model.ReminderID) {
			require.Equal(t, reminder_id_model.ReminderID(123), reminderID)
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
		reminderID: 123,
	}

	reminderRepositoryMock := mock.NewReminderRepositoryMock(mc).
		DeleteReminderMock.
		Inspect(func(ctx context.Context, reminderID reminder_id_model.ReminderID) {
			require.Equal(t, reminder_id_model.ReminderID(123), reminderID)
		}).
		Return(assert.AnError)

	handler := NewHandler(reminderRepositoryMock)

	err := handler.HandleCommand(context.Background(), command)
	require.ErrorIs(t, err, assert.AnError)
}
