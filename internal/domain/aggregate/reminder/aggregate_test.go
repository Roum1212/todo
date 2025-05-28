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

	id, errID := reminder_id_model.NewReminderID("1234")
	title, errTitle := reminder_title_model.NewReminderTitle("title")
	description, errDescription := reminder_description_model.NewReminderDescription("description")

	require.NoError(t, errID, errTitle, errDescription)

	reminder := NewReminder(id, title, description) // я не понимаю, как мне тут удалить reminder и при этом проверить
	// newReminder, хотя в целом тесты для конструкторов можно удалить

	require.Equal(t, id, reminder.GetID())
	require.Equal(t, title, reminder.GetTitle())
	require.Equal(t, description, reminder.GetDescription())
}
