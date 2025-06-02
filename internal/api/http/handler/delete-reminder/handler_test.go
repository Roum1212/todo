package delete_reminder_http_handler

import (
	"context"
	"net/http"
	"net/http/httptest"
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

	commandHandlerMock := mock.NewCommandHandlerMock(mc).
		HandleCommandMock.
		Inspect(func(ctx context.Context, c delete_reminder_command.Command) {
			require.Equal(t, reminder_id_model.ReminderID(123), c.GetReminderID())
		}).
		Return(nil)

	httpHandler := NewHandler(commandHandlerMock)

	router := httprouter.New()
	router.Handler(http.MethodDelete, Endpoint, httpHandler)

	r := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodDelete,
		strings.ReplaceAll(Endpoint, paramsID, "123"),
		http.NoBody,
	)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, r)

	require.Equal(t, http.StatusOK, recorder.Code)
}

func TestHandler_ServeHTTP_BadRequest(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	commandHandlerMock := mock.NewCommandHandlerMock(mc)

	httpHandler := NewHandler(commandHandlerMock)

	router := httprouter.New()
	router.Handler(http.MethodDelete, Endpoint, httpHandler)

	r := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodDelete,
		strings.ReplaceAll(Endpoint, paramsID, "abc"),
		http.NoBody,
	)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, r)

	require.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestHandler_ServeHTTP_InternalServerError(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	commandHandlerMock := mock.NewCommandHandlerMock(mc).
		HandleCommandMock.
		Inspect(func(ctx context.Context, c delete_reminder_command.Command) {
			require.Equal(t, reminder_id_model.ReminderID(123), c.GetReminderID())
		}).
		Return(assert.AnError)

	httpHandler := NewHandler(commandHandlerMock)

	router := httprouter.New()
	router.Handler(http.MethodDelete, Endpoint, httpHandler)

	r := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodDelete,
		strings.ReplaceAll(Endpoint, paramsID, "123"),
		http.NoBody,
	)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, r)

	require.Equal(t, http.StatusInternalServerError, recorder.Code)
}
