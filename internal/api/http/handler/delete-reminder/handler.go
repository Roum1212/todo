package delete_reminder_http_handler

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"

	delete_reminder_command "github.com/Roum1212/todo/internal/app/command/delete-reminder"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

const Endpoint = "/reminders/" + paramID

const paramID = ":id"

type Handler struct {
	commandHandler delete_reminder_command.CommandHandler
}

func (x Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	reminderID, err := reminder_id_model.NewReminderIDFromString(
		params.ByName(strings.TrimPrefix(paramID, ":")),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	if err = x.commandHandler.HandleCommand(
		r.Context(),
		delete_reminder_command.NewCommand(reminderID),
	); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func NewHTTPHandler(commandHandler delete_reminder_command.CommandHandler) Handler {
	return Handler{
		commandHandler: commandHandler,
	}
}
