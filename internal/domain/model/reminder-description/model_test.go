package reminder_description_model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const str = "abc123"

func TestNewReminderDescription(t *testing.T) {
	t.Parallel()

	reminderDescription, err := NewReminderDescription(str)
	require.NoError(t, err)
	require.Equal(t, str, string(reminderDescription))
}

func TestReminderDescription_Validate(t *testing.T) {
	t.Parallel()

	reminderDescription := ReminderDescription(str)
	err := reminderDescription.Validate()
	require.NoError(t, err)
}

func TestReminderDescription_Validate_Error(t *testing.T) {
	t.Parallel()

	reminderDescription := ReminderDescription("")
	err := reminderDescription.Validate()
	require.Error(t, err)
}
