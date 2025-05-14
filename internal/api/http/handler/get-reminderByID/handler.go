package get_reminderByID_http_handler

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	get_reminderByID_quary "github.com/Roum1212/todo/internal/app/quary/get-reminderByID"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

const Endpoint = "/reminders/:id"

type Handler struct {
	commandHandler get_reminderByID_quary.Handler
}

func (x Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id := params.ByName("id")

	reminderID, err := reminder_id_model.NewReminderID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	reminder, err := x.commandHandler.Handle(r.Context(),
		get_reminderByID_quary.NewCommand(reminderID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(reminder); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

}

func NewHandler(commandGetByIDHandler get_reminderByID_quary.Handler) Handler {
	return Handler{
		commandHandler: commandGetByIDHandler,
	}
}
