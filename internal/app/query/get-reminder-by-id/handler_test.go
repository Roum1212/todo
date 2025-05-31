package get_reminder_by_id_query

import (
	"context"
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

const id = 123

func TestQueryHandler_HandleQuery(t *testing.T) {
	t.Parallel()

	t.Run("empty reminder", func(t *testing.T) {
		t.Parallel()

		mc := minimock.NewController(t)

		query := Query{
			reminderID: id,
		}

		reminderRepositoryMock := mock.NewReminderRepositoryMock(mc).
			GetReminderByIDMock.
			Inspect(func(ctx context.Context, reminderID reminder_id_model.ReminderID) {
				require.Equal(t, query.reminderID, reminderID)
			}).
			Return(reminder_aggregate.Reminder{}, nil)

		handler := NewHandler(reminderRepositoryMock)

		reminder, err := handler.HandleQuery(t.Context(), query)
		require.NoError(t, err)
		require.Equal(t, reminder_aggregate.Reminder{}, reminder)
	})

	t.Run("immutable values", func(t *testing.T) {
		t.Parallel()

		mc := minimock.NewController(t)

		query := Query{
			reminderID: id,
		}

		reminderID := reminder_id_model.ReminderID(id)
		reminderTitle := reminder_title_model.ReminderTitle("title")
		reminderDescription := reminder_description_model.ReminderDescription("description")
		reminder := reminder_aggregate.NewReminder(reminderID, reminderTitle, reminderDescription)

		reminderRepositoryMock := mock.NewReminderRepositoryMock(mc).
			GetReminderByIDMock.
			Inspect(func(ctx context.Context, reminderID reminder_id_model.ReminderID) {
				require.Equal(t, query.reminderID, reminderID)
			}).
			Return(reminder, nil)

		handler := NewHandler(reminderRepositoryMock)

		reminderQuery, err := handler.HandleQuery(t.Context(), query)
		require.NoError(t, err)

		require.Equal(t, reminder.GetID(), reminderQuery.GetID())
		require.Equal(t, reminder.GetTitle(), reminderQuery.GetTitle())
		require.Equal(t, reminder.GetDescription(), reminderQuery.GetDescription())
	})
}

func TestQueryHandler_HandleQuery_Err(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	query := Query{
		reminderID: id,
	}

	reminderRepositoryMock := mock.NewReminderRepositoryMock(mc).
		GetReminderByIDMock.
		Inspect(func(ctx context.Context, reminderID reminder_id_model.ReminderID) {
			require.Equal(t, query.reminderID, reminderID)
		}).
		Return(reminder_aggregate.Reminder{}, assert.AnError)

	handler := NewHandler(reminderRepositoryMock)

	reminder, err := handler.HandleQuery(t.Context(), query)
	require.Error(t, err)
	require.Equal(t, reminder_aggregate.Reminder{}, reminder)
}

func TestQueryHandler_HandlerQuery_ErrReminderNotFound(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	query := Query{
		reminderID: id,
	}

	reminderRepositoryMock := mock.NewReminderRepositoryMock(mc).
		GetReminderByIDMock.
		Inspect(func(ctx context.Context, reminderID reminder_id_model.ReminderID) {
			require.Equal(t, query.reminderID, reminderID)
		}).
		Return(reminder_aggregate.Reminder{}, reminder_aggregate.ErrRemindersNotFound)

	handler := NewHandler(reminderRepositoryMock)

	reminder, err := handler.HandleQuery(t.Context(), query)
	require.ErrorIs(t, err, ErrReminderNotFound)
	require.Equal(t, reminder_aggregate.Reminder{}, reminder)
}
