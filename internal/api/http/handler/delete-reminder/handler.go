package delete_reminder_http_handler

import (
	"encoding/json"
	"net/http"

	delete_reminder_command "github.com/Roum1212/todo/internal/app/command/delete-reminder"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

const Endpoint2 = "/reminders/delete"

type HandlerDelete struct {
	commandDeleteHandler delete_reminder_command.Handler
}

func (x HandlerDelete) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request RequestToDelete

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	id := reminder_id_model.NewReminderID(request.ID)

	if err := x.commandDeleteHandler.Handle(
		r.Context(),
		delete_reminder_command.NewDeleteCommand(id),
	); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)
}

func NewHandler(commandDeleteHandler delete_reminder_command.Handler) HandlerDelete {
	return HandlerDelete{
		commandDeleteHandler: commandDeleteHandler,
	}
}

type RequestToDelete struct {
	ID int `json:"id"`
}
