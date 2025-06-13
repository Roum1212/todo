package create_reminder_command

import (
	"crypto/rand"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
	"github.com/Roum1212/todo/internal/domain/aggregate/reminder/mock"
	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
)

func TestCommandHandler_HandleCommand(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	title, err := reminder_title_model.NewReminderTitle(rand.Text())
	require.NoError(t, err)

	description, err := reminder_description_model.NewReminderDescription(rand.Text())
	require.NoError(t, err)

	id := reminder_id_model.GenerateReminderID()

	reminderRepositoryMock := reminder_aggregate_mock.NewReminderRepositoryMock(mc).
		SaveReminderMock.
		Expect(minimock.AnyContext, reminder_aggregate.NewReminder(id, title, description)).
		Return(nil)

	handler := NewCommandHandler(reminderRepositoryMock)
	require.NoError(t, handler.HandleCommand(t.Context(), NewCommand(id, title, description)))
}

func TestCommandHandler_HandleCommand_Error(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	title, err := reminder_title_model.NewReminderTitle(rand.Text())
	require.NoError(t, err)

	description, err := reminder_description_model.NewReminderDescription(rand.Text())
	require.NoError(t, err)

	id := reminder_id_model.GenerateReminderID()

	reminderRepositoryMock := reminder_aggregate_mock.NewReminderRepositoryMock(mc).
		SaveReminderMock.
		Expect(minimock.AnyContext, reminder_aggregate.NewReminder(id, title, description)).
		Return(assert.AnError)

	handler := NewCommandHandler(reminderRepositoryMock)
	require.ErrorIs(t, handler.HandleCommand(t.Context(), NewCommand(id, title, description)), assert.AnError)
}

func TestTracerCommandHandler_HandleCommand(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	title, err := reminder_title_model.NewReminderTitle(rand.Text())
	require.NoError(t, err)

	description, err := reminder_description_model.NewReminderDescription(rand.Text())
	require.NoError(t, err)

	id := reminder_id_model.GenerateReminderID()

	command := NewCommand(id, title, description)

	reminderRepositoryMock := reminder_aggregate_mock.NewReminderRepositoryMock(mc).
		SaveReminderMock.
		Expect(minimock.AnyContext, reminder_aggregate.NewReminder(id, title, description)).
		Return(nil)

	handler := NewCommandHandler(reminderRepositoryMock)
	require.NoError(t, handler.HandleCommand(t.Context(), command))

	handlerTracer := NewCommandHandlerWithTracing(handler)
	require.NoError(t, handlerTracer.HandleCommand(t.Context(), command))
}

func TestTracerCommandHandler_HandleCommand_Error(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	title, err := reminder_title_model.NewReminderTitle(rand.Text())
	require.NoError(t, err)

	description, err := reminder_description_model.NewReminderDescription(rand.Text())
	require.NoError(t, err)

	id := reminder_id_model.GenerateReminderID()

	command := NewCommand(id, title, description)

	reminderRepositoryMock := reminder_aggregate_mock.NewReminderRepositoryMock(mc).
		SaveReminderMock.
		Expect(minimock.AnyContext, reminder_aggregate.NewReminder(id, title, description)).
		Return(assert.AnError)

	handler := NewCommandHandler(reminderRepositoryMock)
	require.Error(t, handler.HandleCommand(t.Context(), command))

	handlerTracer := NewCommandHandlerWithTracing(handler)
	require.Error(t, handlerTracer.HandleCommand(t.Context(), command))
}
