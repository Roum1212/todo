package postgresql_reminder_repository

import (
	"testing"

	"github.com/stretchr/testify/require"

	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
)

// func NewReminders(reminderDTOs []Reminder) ([]reminder_aggregate.Reminder, error)

func TestNewReminders(t *testing.T) {
	t.Parallel()

	t.Run("nil slice", func(t *testing.T) {
		t.Parallel()

		reminders, err := NewReminders(nil)
		require.NoError(t, err)
		require.Len(t, reminders, 0)
	})

	t.Run("empty slice", func(t *testing.T) {
		t.Parallel()

		reminders, err := NewReminders([]Reminder{})
		require.NoError(t, err)
		require.Len(t, reminders, 0)
	})

	t.Run("immutable values", func(t *testing.T) {
		tests := []Reminder{
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
			{
				ID:          789,
				Title:       "titleT",
				Description: "descriptionD",
			},
		}

		reminders, err := NewReminders(tests)
		require.NoError(t, err)
		require.Len(t, reminders, 3)

		for i := range reminders {
			require.Equal(
				t,
				reminders[i].GetID(),
				reminder_id_model.ReminderID(tests[i].ID),
			)
			require.Equal(
				t,
				reminders[i].GetTitle(),
				reminder_title_model.ReminderTitle(tests[i].Title),
			)
			require.Equal(
				t,
				reminders[i].GetDescription(),
				reminder_description_model.ReminderDescription(tests[i].Description),
			)
		}
	})

}
