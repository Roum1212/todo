package create_reminder_command

import (
	"context"
	"crypto/rand"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
	"github.com/Roum1212/todo/internal/domain/aggregate/reminder/mock"
	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
)

func TestCommandHandler_HandleCommand(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	title, err := reminder_title_model.NewReminderTitle(rand.Text())
	require.NoError(t, err)

	description, err := reminder_description_model.NewReminderDescription(rand.Text())
	require.NoError(t, err)

	command := NewCommand(title, description)

	reminderRepositoryMock := reminder_aggregate_mock.NewReminderRepositoryMock(mc).
		SaveReminderMock.
		Inspect(func(ctx context.Context, reminder reminder_aggregate.Reminder) {
			require.Equal(t, title, reminder.GetTitle())
			require.Equal(t, description, reminder.GetDescription())
		}).
		Return(nil)

	handler := NewCommandHandler(reminderRepositoryMock)
	require.NoError(t, handler.HandleCommand(t.Context(), command))
}

func TestCommandHandler_HandleCommand_Error(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	title, err := reminder_title_model.NewReminderTitle(rand.Text())
	require.NoError(t, err)

	description, err := reminder_description_model.NewReminderDescription(rand.Text())
	require.NoError(t, err)

	command := NewCommand(title, description)

	reminderRepositoryMock := reminder_aggregate_mock.NewReminderRepositoryMock(mc).
		SaveReminderMock.
		Inspect(func(ctx context.Context, reminder reminder_aggregate.Reminder) {
			require.Equal(t, title, reminder.GetTitle())
			require.Equal(t, description, reminder.GetDescription())
		}).
		Return(assert.AnError)

	handler := NewCommandHandler(reminderRepositoryMock)
	require.ErrorIs(t, handler.HandleCommand(t.Context(), command), assert.AnError)
}

func TestTracerCommandHandler_HandleCommand(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	title, err := reminder_title_model.NewReminderTitle(rand.Text())
	require.NoError(t, err)

	description, err := reminder_description_model.NewReminderDescription(rand.Text())
	require.NoError(t, err)

	command := NewCommand(title, description)

	reminderRepositoryMock := reminder_aggregate_mock.NewReminderRepositoryMock(mc).
		SaveReminderMock.
		Inspect(func(ctx context.Context, reminder reminder_aggregate.Reminder) {
			require.Equal(t, title, reminder.GetTitle())
			require.Equal(t, description, reminder.GetDescription())
		}).
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

	command := NewCommand(title, description)

	reminderRepositoryMock := reminder_aggregate_mock.NewReminderRepositoryMock(mc).
		SaveReminderMock.
		Inspect(func(ctx context.Context, reminder reminder_aggregate.Reminder) {
			require.Equal(t, title, reminder.GetTitle())
			require.Equal(t, description, reminder.GetDescription())
		}).
		Return(assert.AnError)

	handler := NewCommandHandler(reminderRepositoryMock)
	require.Error(t, handler.HandleCommand(t.Context(), command))

	handlerTracer := NewCommandHandlerWithTracing(handler)
	require.Error(t, handlerTracer.HandleCommand(t.Context(), command))
}
