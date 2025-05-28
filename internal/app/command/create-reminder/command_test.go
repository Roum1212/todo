package create_reminder_command

import (
	"testing"

	"github.com/stretchr/testify/require"

	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
)

func TestNewCommand(t *testing.T) {
	t.Parallel()

	title, errTitle := reminder_title_model.NewReminderTitle("title")
	description, errDescription := reminder_description_model.NewReminderDescription("description")

	require.NoError(t, errTitle, errDescription)

	command := NewCommand(title, description)

	require.Equal(t, title, command.title)
	require.Equal(t, description, command.description)
}
