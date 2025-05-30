package reminder_title_model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const str = "abc123"

func TestNewReminderTitle(t *testing.T) {
	t.Parallel()

	reminderTitle, err := NewReminderTitle(str)
	require.NoError(t, err)
	require.Equal(t, str, string(reminderTitle))
}

func TestReminderTitle_Validate(t *testing.T) {
	t.Parallel()

	reminderTitle := ReminderTitle(str)
	err := reminderTitle.Validate()
	require.NoError(t, err)
}

func TestReminderTitle_Validate_Error(t *testing.T) {
	t.Parallel()

	reminderTitle := ReminderTitle("")
	err := reminderTitle.Validate()
	require.Error(t, err)
}
