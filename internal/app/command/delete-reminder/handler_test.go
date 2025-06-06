package delete_reminder_command

import (
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Roum1212/todo/internal/domain/aggregate/reminder/mock"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

func TestCommandHandler_HandleCommand(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminderID := reminder_id_model.GenerateReminderID()

	command := NewCommand(reminderID)

	reminderRepositoryMock := reminder_aggregate_mock.NewReminderRepositoryMock(mc).
		DeleteReminderMock.
		Expect(minimock.AnyContext, reminderID).
		Return(nil)

	handler := NewCommandHandler(reminderRepositoryMock)
	require.NoError(t, handler.HandleCommand(t.Context(), command))
}

func TestCommandHandler_HandleCommand_Error(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminderID := reminder_id_model.GenerateReminderID()

	command := NewCommand(reminderID)

	reminderRepositoryMock := reminder_aggregate_mock.NewReminderRepositoryMock(mc).
		DeleteReminderMock.
		Expect(minimock.AnyContext, reminderID).
		Return(assert.AnError)

	handler := NewCommandHandler(reminderRepositoryMock)
	require.ErrorIs(t, handler.HandleCommand(t.Context(), command), assert.AnError)
}
