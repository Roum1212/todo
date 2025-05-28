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

	TableTest := []struct {
		id          string
		title       string
		description string
	}{
		{id: "1", title: "title", description: "description"},
		{id: "-10", title: "", description: "description"},
		{id: "0", title: "title", description: ""},
	}

	for _, tt := range TableTest {
		id, _ := reminder_id_model.NewReminderID(tt.id) //nolint:errcheck // OK.
		title, _ := reminder_title_model.NewReminderTitle(tt.title)
		description, _ := reminder_description_model.NewReminderDescription(tt.description)

		reminder := NewReminder(id, title, description)

		require.Equal(t, id, reminder.GetID())
		require.Equal(t, title, reminder.GetTitle())
		require.Equal(t, description, reminder.GetDescription())
	}
}
