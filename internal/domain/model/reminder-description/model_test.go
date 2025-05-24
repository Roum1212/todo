package reminder_description_model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewReminderDescription(t *testing.T) {
	t.Parallel()
	reminderDescription := NewReminderDescription("abc123") //nolint:wsl // OK.

	require.Equal(t, ReminderDescription("abc123"), reminderDescription)
}
