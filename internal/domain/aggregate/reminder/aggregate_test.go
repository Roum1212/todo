package reminder_aggregate

import (
	"testing"

	"github.com/stretchr/testify/require"

	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
)

func TestNewReminder(t *testing.T) {
	t.Parallel()

	id, err := reminder_id_model.NewReminderID("123")
	require.NoError(t, err)

	title, err := reminder_title_model.NewReminderTitle("abc")
	require.NoError(t, err)

	description, err := reminder_description_model.NewReminderDescription("abc")
	require.NoError(t, err)

	reminder := NewReminder(id, title, description)
	require.Equal(t, id, reminder.GetID())
	require.Equal(t, title, reminder.GetTitle())
	require.Equal(t, description, reminder.GetDescription())
}
