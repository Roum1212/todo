package reminder_description_model

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewReminderDescription(t *testing.T) {
	t.Parallel()

	s := rand.Text()

	reminderDescription, err := NewReminderDescription(s)
	require.NoError(t, err)
	require.Equal(t, s, string(reminderDescription))
}

func TestReminderDescription_Validate(t *testing.T) {
	t.Parallel()

	s := rand.Text()

	reminderDescription, err := NewReminderDescription(s)
	require.NoError(t, err)
	require.NoError(t, reminderDescription.Validate())
}

func TestReminderDescription_Validate_Error(t *testing.T) {
	t.Parallel()

	reminderDescription := ReminderDescription("")
	require.Error(t, reminderDescription.Validate())
}
