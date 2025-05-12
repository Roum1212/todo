package delete_reminder_http_handler

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	delete_reminder_command "github.com/Roum1212/todo/internal/app/command/delete-reminder"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

const Endpoint = "/reminders/:id"

type Handler struct {
	commandDeleteHandler delete_reminder_command.Handler
}

func (x Handler) ServeHTTP(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "ID must be a number", http.StatusBadRequest)
		return
	}

	if err := x.commandDeleteHandler.Handle(
		r.Context(),
		delete_reminder_command.NewDeleteCommand(reminder_id_model.ReminderID(idInt)),
	); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)
}

func NewHandler(commandDeleteHandler delete_reminder_command.Handler) Handler {
	return Handler{
		commandDeleteHandler: commandDeleteHandler,
	}
}
