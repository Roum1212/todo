package delete_reminder_http_handler

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"

	delete_reminder_command "github.com/Roum1212/todo/internal/app/command/delete-reminder"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

const paramsID = ":id"

const Endpoint = "/reminders/" + paramsID

type Handler struct {
	commandHandler delete_reminder_command.CommandHandler
}

func (x Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id := params.ByName(strings.TrimPrefix(paramsID, ":"))

	reminderID, err := reminder_id_model.NewReminderID(id)
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

func NewHandler(commandDeleteHandler delete_reminder_command.CommandHandler) Handler {
	return Handler{
		commandHandler: commandDeleteHandler,
	}
}
