package get_all_reminders_http_handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	get_all_reminders_quer "github.com/Roum1212/todo/internal/app/query/get-all-reminders"
	"github.com/Roum1212/todo/internal/app/query/get-all-reminders/mock"
	postgresql_reminder_repository "github.com/Roum1212/todo/internal/infra/repository/reminder/postgresql"
)

func TestHandler_ServeHTTP_OK(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminderDTOs := []postgresql_reminder_repository.Reminder{
		{
			ID:          123,
			Title:       "Title",
			Description: "Description",
		},
		{
			ID:          456,
			Title:       "title",
			Description: "description",
		},
		{
			ID:          789,
			Title:       "titleT",
			Description: "descriptionD",
		},
	}

	reminders, err := postgresql_reminder_repository.NewReminders(reminderDTOs)
	require.NoError(t, err)

	queryHandlerMock := mock.NewQueryHandlerMock(mc).
		HandleQueryMock.
		Return(reminders, nil)

	httpHandler := NewHandler(queryHandlerMock)

	r := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodGet,
		Endpoint,
		http.NoBody)

	recorder := httptest.NewRecorder()

	httpHandler.ServeHTTP(recorder, r)

	require.Equal(t, http.StatusOK, recorder.Code)
	require.Equal(t, "application/json", recorder.Header().Get("Content-Type"))

	expectedBody := NewReminderDTOs(reminders)

	var gotBody []Reminder
	err = json.NewDecoder(recorder.Body).Decode(&gotBody)
	require.NoError(t, err)
	require.Equal(t, expectedBody, gotBody)
}

func TestHandler_ServeHTTP_InternalServerError(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	queryHandlerMock := mock.NewQueryHandlerMock(mc).
		HandleQueryMock.
		Return(nil, assert.AnError)

	httpHandler := NewHandler(queryHandlerMock)

	r := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodGet,
		Endpoint,
		http.NoBody)

	recorder := httptest.NewRecorder()

	httpHandler.ServeHTTP(recorder, r)

	require.Equal(t, http.StatusInternalServerError, recorder.Code)
}

func TestHandler_ServeHTTP_StatusNotFound(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	queryHandlerMock := mock.NewQueryHandlerMock(mc).
		HandleQueryMock.
		Return(nil, get_all_reminders_quer.ErrRemindersNotFound)

	httpHandler := NewHandler(queryHandlerMock)

	r := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodGet,
		Endpoint,
		http.NoBody)

	recorder := httptest.NewRecorder()

	httpHandler.ServeHTTP(recorder, r)

	require.Equal(t, http.StatusNotFound, recorder.Code)
}
