package get_all_reminders_query

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

	reminder1 := reminder_aggregate.NewReminder(
		reminder_id_model.GenerateReminderID(),
		reminder_title_model.ReminderTitle(rand.Text()),
		reminder_description_model.ReminderDescription(rand.Text()),
	)
	reminder2 := reminder_aggregate.NewReminder(
		reminder_id_model.GenerateReminderID(),
		reminder_title_model.ReminderTitle(rand.Text()),
		reminder_description_model.ReminderDescription(rand.Text()),
	)
	reminders := []reminder_aggregate.Reminder{reminder1, reminder2}

	reminderRepositoryMock := reminder_aggregate_mock.NewReminderRepositoryMock(mc).
		GetAllRemindersMock.
		Expect(minimock.AnyContext).
		Return(reminders, nil)

	handler := NewQueryHandler(reminderRepositoryMock)

	gotReminders, err := handler.HandleQuery(t.Context())
	require.NoError(t, err)
	require.Equal(t, reminders, gotReminders)
}

func TestQueryHandler_HandleQuery_Error(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminderRepositoryMock := reminder_aggregate_mock.NewReminderRepositoryMock(mc).
		GetAllRemindersMock.
		Expect(minimock.AnyContext).
		Return(nil, assert.AnError)

	handler := NewQueryHandler(reminderRepositoryMock)

	gotReminders, err := handler.HandleQuery(t.Context())
	require.Error(t, err)
	require.Empty(t, gotReminders)
}

func TestQueryHandler_HandleQuery_ErrRemindersNotFound(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminderRepositoryMock := reminder_aggregate_mock.NewReminderRepositoryMock(mc).
		GetAllRemindersMock.
		Expect(minimock.AnyContext).
		Return(nil, reminder_aggregate.ErrReminderNotFound)

	handler := NewQueryHandler(reminderRepositoryMock)

	gotReminders, err := handler.HandleQuery(t.Context())
	require.ErrorIs(t, err, ErrReminderNotFound)
	require.Empty(t, gotReminders)
}
