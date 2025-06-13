package delete_reminder_command

import (
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
	"github.com/Roum1212/todo/internal/domain/aggregate/reminder/mock"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

func TestCommandHandler_HandleCommand(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminderID := reminder_id_model.GenerateReminderID()

	reminderRepositoryMock := reminder_aggregate_mock.NewReminderRepositoryMock(mc).
		DeleteReminderMock.
		Expect(minimock.AnyContext, reminderID).
		Return(nil)

	handler := NewCommandHandler(reminderRepositoryMock)
	require.NoError(t, handler.HandleCommand(t.Context(), NewCommand(reminderID)))
}

func TestCommandHandler_HandleCommand_Error(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminderID := reminder_id_model.GenerateReminderID()

	reminderRepositoryMock := reminder_aggregate_mock.NewReminderRepositoryMock(mc).
		DeleteReminderMock.
		Expect(minimock.AnyContext, reminderID).
		Return(assert.AnError)

	handler := NewCommandHandler(reminderRepositoryMock)
	require.ErrorIs(t, handler.HandleCommand(t.Context(), NewCommand(reminderID)), assert.AnError)
}

func TestCommandHandler_HandleCommand_ErrReminderNotFound(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminderID := reminder_id_model.GenerateReminderID()

	reminderRepositoryMock := reminder_aggregate_mock.NewReminderRepositoryMock(mc).
		DeleteReminderMock.
		Expect(t.Context(), reminderID).
		Return(reminder_aggregate.ErrReminderNotFound)

	handler := NewCommandHandler(reminderRepositoryMock)
	require.ErrorIs(t, handler.HandleCommand(t.Context(), NewCommand(reminderID)), ErrReminderNotFound)
}

func TestTracerCommandHandler_HandleCommand(t *testing.T) {
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

	handlerTracer := NewCommandHandlerTracer(handler)
	require.NoError(t, handlerTracer.HandleCommand(t.Context(), command))
}

func TestTracerCommandHandler_HandleCommand_Error(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminderID := reminder_id_model.GenerateReminderID()

	command := NewCommand(reminderID)

	reminderRepositoryMock := reminder_aggregate_mock.NewReminderRepositoryMock(mc).
		DeleteReminderMock.
		Expect(minimock.AnyContext, reminderID).
		Return(assert.AnError)

	handler := NewCommandHandler(reminderRepositoryMock)
	require.Error(t, handler.HandleCommand(t.Context(), command))

	handlerTracer := NewCommandHandlerTracer(handler)
	require.Error(t, handlerTracer.HandleCommand(t.Context(), command))
}
