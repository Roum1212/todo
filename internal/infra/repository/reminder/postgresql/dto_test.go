package postgresql_reminder_repository

import (
	"testing"

	"github.com/stretchr/testify/require"

	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
)

func TestNewReminders_Success(t *testing.T) {
	t.Parallel()

	test := []Reminder{
		{
			ID: 123, Title: "title", Description: "description",
		},
	}

	reminder, err := NewReminders(test)
	require.NoError(t, err)

	require.Equal(t, reminder_id_model.ReminderID(123), reminder[0].GetID())
	require.Equal(t, reminder_title_model.ReminderTitle("title"), reminder[0].GetTitle())
	require.Equal(t, reminder_description_model.ReminderDescription("description"), reminder[0].GetDescription())
}

func TestNewReminders_Error(t *testing.T) {
	t.Parallel()

	test := []Reminder{
		{
			ID: 123, Title: "", Description: "description",
		},
		{
			ID: 123, Title: "title", Description: "",
		},
	}

	for _, tt := range test {
		_, err := NewReminders([]Reminder{tt})
		require.Error(t, err)
	}
}
