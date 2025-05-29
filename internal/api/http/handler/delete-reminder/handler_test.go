package delete_reminder_http_handler

import (
	"context"
	"net/http"
	"net/http/httptest"
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

	const id = "123"

	commandHandlerMock := mock.NewCommandHandlerMock(mc).
		HandleCommandMock.
		Inspect(func(ctx context.Context, c delete_reminder_command.Command) {
			require.Equal(t, reminder_id_model.ReminderID(123), c.GetID())
		}).
		Return(nil)

	httpHandler := NewHandler(commandHandlerMock)

	r := httptest.NewRequest(http.MethodDelete, "/reminders/"+id, nil)

	params := httprouter.Params{httprouter.Param{Key: "id", Value: id}}
	ctx := context.WithValue(r.Context(), httprouter.ParamsKey, params)
	r = r.WithContext(ctx)

	recorder := httptest.NewRecorder()

	httpHandler.ServeHTTP(recorder, r)

	require.Equal(t, http.StatusOK, recorder.Code)
}

func TestHandler_ServeHTTP_BadRequest(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	const id = ""

	commandHandlerMock := mock.NewCommandHandlerMock(mc)

	httpHandler := NewHandler(commandHandlerMock)

	r := httptest.NewRequest(http.MethodDelete, "/reminders/"+id, nil)

	params := httprouter.Params{httprouter.Param{Key: "id", Value: id}}
	ctx := context.WithValue(r.Context(), httprouter.ParamsKey, params)
	r = r.WithContext(ctx)

	recorder := httptest.NewRecorder()

	httpHandler.ServeHTTP(recorder, r)

	require.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestHandler_ServeHTTP_InternalServerError(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	const id = "123"

	commandHandlerMock := mock.NewCommandHandlerMock(mc).
		HandleCommandMock.
		Inspect(func(ctx context.Context, c delete_reminder_command.Command) {
			require.Equal(t, reminder_id_model.ReminderID(123), c.GetID())
		}).
		Return(assert.AnError)

	httpHandler := NewHandler(commandHandlerMock)

	r := httptest.NewRequest(http.MethodDelete, "/reminders/"+id, nil)

	params := httprouter.Params{httprouter.Param{Key: "id", Value: id}}
	ctx := context.WithValue(r.Context(), httprouter.ParamsKey, params)
	r = r.WithContext(ctx)

	recorder := httptest.NewRecorder()

	httpHandler.ServeHTTP(recorder, r)

	require.Equal(t, http.StatusInternalServerError, recorder.Code)
}
