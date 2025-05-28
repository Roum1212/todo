package reminder_description_model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewReminderDescription(t *testing.T) {
	t.Parallel()

	reminderDescription, _ := NewReminderDescription("abc123")
	require.Equal(t, "abc123", string(reminderDescription))
}
