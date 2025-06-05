package delete_reminder_http_handler

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	delete_reminder_command "github.com/Roum1212/todo/internal/app/command/delete-reminder"
	"github.com/Roum1212/todo/internal/app/command/delete-reminder/mock"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

func TestHandler_ServeHTTP_OK(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminderID := reminder_id_model.GenerateReminderID()

	commandHandlerMock := delete_reminder_command_mock.NewCommandHandlerMock(mc).
		HandleCommandMock.
		Expect(minimock.AnyContext, delete_reminder_command.NewCommand(reminderID)).
		Return(nil)

	httpHandler := NewHTTPHandler(commandHandlerMock)

	router := httprouter.New()
	router.Handler(http.MethodDelete, Endpoint, httpHandler)

	r := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodDelete,
		strings.ReplaceAll(Endpoint, paramID, strconv.FormatInt(int64(reminderID), 10)),
		http.NoBody,
	)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, r)

	require.Equal(t, http.StatusOK, recorder.Code)
}

func TestHandler_ServeHTTP_BadRequest(t *testing.T) {
	t.Parallel()

	httpHandler := NewHTTPHandler(nil)

	router := httprouter.New()
	router.Handler(http.MethodDelete, Endpoint, httpHandler)

	r := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodDelete,
		strings.ReplaceAll(Endpoint, paramID, "abc"),
		http.NoBody,
	)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, r)

	require.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestHandler_ServeHTTP_InternalServerError(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminderID := reminder_id_model.GenerateReminderID()

	commandHandlerMock := delete_reminder_command_mock.NewCommandHandlerMock(mc).
		HandleCommandMock.
		Expect(minimock.AnyContext, delete_reminder_command.NewCommand(reminderID)).
		Return(assert.AnError)

	httpHandler := NewHTTPHandler(commandHandlerMock)

	router := httprouter.New()
	router.Handler(http.MethodDelete, Endpoint, httpHandler)

	r := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodDelete,
		strings.ReplaceAll(Endpoint, paramID, strconv.FormatInt(int64(reminderID), 10)),
		http.NoBody,
	)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, r)

	require.Equal(t, http.StatusInternalServerError, recorder.Code)
}
