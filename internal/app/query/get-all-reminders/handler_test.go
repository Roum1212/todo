package get_all_reminders_query

import (
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
	"github.com/Roum1212/todo/internal/domain/aggregate/reminder/mock"
	postgresql_reminder_repository "github.com/Roum1212/todo/internal/infra/repository/reminder/postgresql"
)

func TestQueryHandler_HandleQuery(t *testing.T) {
	t.Parallel()

	t.Run("immutable values", func(t *testing.T) {
		t.Parallel()

		mc := minimock.NewController(t)

		reminderDTOs := []postgresql_reminder_repository.Reminder{
			{
				ID:          123,
				Title:       "title",
				Description: "description",
			},
			{
				ID:          456,
				Title:       "Title",
				Description: "Description",
			},
		}

		reminders, err := postgresql_reminder_repository.NewReminders(reminderDTOs)
		require.NoError(t, err)

		reminderRepositoryMock := mock.NewReminderRepositoryMock(mc).
			GetAllRemindersMock.
			Return(reminders, nil)

		handler := NewHandler(reminderRepositoryMock)

		reminderSlice, err := handler.HandleQuery(t.Context())
		require.NoError(t, err)
		require.Len(t, reminderSlice, len(reminders))

		for i := range reminderSlice {
			require.Equal(t, reminders[i].GetID(), reminderSlice[i].GetID())
			require.Equal(t, reminders[i].GetTitle(), reminderSlice[i].GetTitle())
			require.Equal(t, reminders[i].GetDescription(), reminderSlice[i].GetDescription())
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		t.Parallel()

		mc := minimock.NewController(t)

		reminderRepositoryMock := mock.NewReminderRepositoryMock(mc).
			GetAllRemindersMock.
			Return(nil, nil)

		handler := NewHandler(reminderRepositoryMock)

		reminderSlice, err := handler.HandleQuery(t.Context())
		require.NoError(t, err)
		require.Nil(t, reminderSlice)
		require.Len(t, reminderSlice, 0)
	})
}

func TestQueryHandler_HandleQuery_Error(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminderRepositoryMock := mock.NewReminderRepositoryMock(mc).
		GetAllRemindersMock.
		Return(nil, assert.AnError)

	handler := NewHandler(reminderRepositoryMock)

	reminderSlice, err := handler.HandleQuery(t.Context())
	require.Error(t, err)
	require.Nil(t, reminderSlice)
	require.Len(t, reminderSlice, 0)
}

func TestQueryHandler_HandleQuery_ErrRemindersNotFound(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminderRepositoryMock := mock.NewReminderRepositoryMock(mc).
		GetAllRemindersMock.
		Return(nil, reminder_aggregate.ErrRemindersNotFound)

	handler := NewHandler(reminderRepositoryMock)

	reminderSlice, err := handler.HandleQuery(t.Context())
	require.ErrorIs(t, err, ErrRemindersNotFound)
	require.Nil(t, reminderSlice)
}
