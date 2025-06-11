package get_reminder_by_id_query

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

func TestQueryHandler_HandleQuery(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminderID := reminder_id_model.GenerateReminderID()
	reminder := reminder_aggregate.NewReminder(
		reminderID,
		reminder_title_model.ReminderTitle(rand.Text()),
		reminder_description_model.ReminderDescription(rand.Text()),
	)

	query := NewQuery(reminderID)

	reminderRepositoryMock := reminder_aggregate_mock.NewReminderRepositoryMock(mc).
		GetReminderByIDMock.
		Expect(minimock.AnyContext, reminderID).
		Return(reminder, nil)

	handler := NewQueryHandler(reminderRepositoryMock)

	gotReminder, err := handler.HandleQuery(t.Context(), query)
	require.NoError(t, err)
	require.Equal(t, reminder, gotReminder)
}

func TestQueryHandler_HandleQuery_Err(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminderID := reminder_id_model.GenerateReminderID()

	query := NewQuery(reminderID)

	reminderRepositoryMock := reminder_aggregate_mock.NewReminderRepositoryMock(mc).
		GetReminderByIDMock.
		Expect(minimock.AnyContext, reminderID).
		Return(reminder_aggregate.Reminder{}, assert.AnError)

	handler := NewQueryHandler(reminderRepositoryMock)

	gotReminder, err := handler.HandleQuery(t.Context(), query)
	require.Error(t, err)
	require.Zero(t, gotReminder)
}

func TestQueryHandler_HandlerQuery_ErrReminderNotFound(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminderID := reminder_id_model.GenerateReminderID()

	query := NewQuery(reminderID)

	reminderRepositoryMock := reminder_aggregate_mock.NewReminderRepositoryMock(mc).
		GetReminderByIDMock.
		Expect(minimock.AnyContext, reminderID).
		Return(reminder_aggregate.Reminder{}, reminder_aggregate.ErrReminderNotFound)

	handler := NewQueryHandler(reminderRepositoryMock)

	gotReminder, err := handler.HandleQuery(t.Context(), query)
	require.ErrorIs(t, err, ErrReminderNotFound)
	require.Zero(t, gotReminder)
}

func TestTracerQueryHandler_HandleQuery(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminderID := reminder_id_model.GenerateReminderID()
	reminder := reminder_aggregate.NewReminder(
		reminderID,
		reminder_title_model.ReminderTitle(rand.Text()),
		reminder_description_model.ReminderDescription(rand.Text()),
	)

	query := NewQuery(reminderID)

	reminderRepositoryMock := reminder_aggregate_mock.NewReminderRepositoryMock(mc).
		GetReminderByIDMock.
		Expect(minimock.AnyContext, reminderID).
		Return(reminder, nil)

	handler := NewQueryHandler(reminderRepositoryMock)

	gotReminder, err := handler.HandleQuery(t.Context(), query)
	require.NoError(t, err)
	require.Equal(t, reminder, gotReminder)

	handlerTracer := NewQueryHandlerTracer(handler)
	gotReminder, err = handlerTracer.HandleQuery(t.Context(), query)
	require.NoError(t, err)
	require.Equal(t, reminder, gotReminder)
}

func TestTracerQueryHandler_HandleQuery_Error(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminderID := reminder_id_model.GenerateReminderID()

	query := NewQuery(reminderID)

	reminderRepositoryMock := reminder_aggregate_mock.NewReminderRepositoryMock(mc).
		GetReminderByIDMock.
		Expect(minimock.AnyContext, reminderID).
		Return(reminder_aggregate.Reminder{}, assert.AnError)

	handler := NewQueryHandler(reminderRepositoryMock)

	gotReminder, err := handler.HandleQuery(t.Context(), query)
	require.Error(t, err)
	require.Zero(t, gotReminder)

	handlerTracer := NewQueryHandlerTracer(handler)
	gotReminder, err = handlerTracer.HandleQuery(t.Context(), query)
	require.ErrorIs(t, err, assert.AnError)
	require.Zero(t, gotReminder)
}
