package create_reminder_command

import (
	"testing"

	"github.com/stretchr/testify/require"

	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
)

func TestNewCommand(t *testing.T) {
	t.Parallel()

	TableTest := []struct {
		title, description string
	}{
		{title: "title", description: "description"},
		{title: "title", description: ""},
		{title: "", description: "description"},
	}

	for _, tt := range TableTest {
		title, _ := reminder_title_model.NewReminderTitle(tt.title)
		description, _ := reminder_description_model.NewReminderDescription(tt.description)

		command := NewCommand(title, description)
		require.Equal(t, title, command.title)
		require.Equal(t, description, command.description)
	}
}
