package get_reminder_by_id_http_handler

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	get_reminder_by_id_quary "github.com/Roum1212/todo/internal/app/query/get-reminder-by-id"
	"github.com/Roum1212/todo/internal/app/query/get-reminder-by-id/mock"
	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
)

func TestHandler_ServeHTTP_OK(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminder := reminder_aggregate.NewReminder(
		reminder_id_model.GenerateReminderID(),
		reminder_title_model.ReminderTitle(rand.Text()),
		reminder_description_model.ReminderDescription(rand.Text()),
	)

	queryHandlerMock := mock.NewQueryHandlerMock(mc).
		HandleQueryMock.
		Inspect(func(ctx context.Context, q get_reminder_by_id_quary.Query) {
			require.Equal(t, reminder.GetID(), q.GetReminderID())
		}).
		Return(reminder, nil)

	httpHandler := NewHandler(queryHandlerMock)

	router := httprouter.New()
	router.Handler(http.MethodGet, Endpoint, httpHandler)

	r := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodGet,
		strings.ReplaceAll(Endpoint, paramID, strconv.FormatInt(int64(reminder.GetID()), 10)),
		http.NoBody,
	)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, r)

	require.Equal(t, http.StatusOK, recorder.Code)
	require.Equal(t, "application/json", recorder.Header().Get("Content-Type"))

	reminderDTO := NewReminder(reminder)

	var gotReminderDTO Reminder

	require.NoError(t, json.NewDecoder(recorder.Body).Decode(&gotReminderDTO))
	assert.Equal(t, reminderDTO, gotReminderDTO)
}

func TestHandler_ServeHTTP_BadRequest(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	queryHandlerMock := mock.NewQueryHandlerMock(mc)

	httpHandler := NewHandler(queryHandlerMock)

	router := httprouter.New()
	router.Handler(http.MethodGet, Endpoint, httpHandler)

	r := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodGet,
		strings.ReplaceAll(Endpoint, paramID, "abc"),
		http.NoBody,
	)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, r)

	require.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestHandler_ServeHTTP_StatusNotFound(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminderID := reminder_id_model.GenerateReminderID()

	queryHandlerMock := mock.NewQueryHandlerMock(mc).
		HandleQueryMock.
		Inspect(func(ctx context.Context, q get_reminder_by_id_quary.Query) {
			require.Equal(t, reminderID, q.GetReminderID())
		}).
		Return(reminder_aggregate.Reminder{}, get_reminder_by_id_quary.ErrReminderNotFound)

	httpHandler := NewHandler(queryHandlerMock)

	router := httprouter.New()
	router.Handler(http.MethodGet, Endpoint, httpHandler)

	r := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodGet,
		strings.ReplaceAll(Endpoint, paramID, strconv.FormatInt(int64(reminderID), 10)),
		http.NoBody,
	)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, r)

	require.Equal(t, http.StatusNotFound, recorder.Code)
}

func TestHandler_ServeHTTP_InternalServerError(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminderID := reminder_id_model.GenerateReminderID()

	queryHandlerMock := mock.NewQueryHandlerMock(mc).
		HandleQueryMock.
		Inspect(func(ctx context.Context, q get_reminder_by_id_quary.Query) {
			require.Equal(t, reminderID, q.GetReminderID())
		}).
		Return(reminder_aggregate.Reminder{}, assert.AnError)

	httpHandler := NewHandler(queryHandlerMock)

	router := httprouter.New()
	router.Handler(http.MethodGet, Endpoint, httpHandler)

	r := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodGet,
		strings.ReplaceAll(Endpoint, paramID, strconv.FormatInt(int64(reminderID), 10)),
		http.NoBody,
	)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, r)

	require.Equal(t, http.StatusInternalServerError, recorder.Code)
}
