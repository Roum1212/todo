package reminder_description_model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewReminderDescription(t *testing.T) {
	t.Parallel()

	reminderDescription, err := NewReminderDescription("abc123")

	require.NoError(t, err)
	require.Equal(t, "abc123", string(reminderDescription))
}
