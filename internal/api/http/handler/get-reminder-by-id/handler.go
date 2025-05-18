package get_reminder_by_id_http_handler

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	get_reminderByID_quary "github.com/Roum1212/todo/internal/app/query/get-reminder-by-id"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	postgresql_reminder_repository "github.com/Roum1212/todo/internal/infra/repository/reminder/postgresql"
)

const Endpoint = "/reminders/:id"

type Handler struct {
	queryHandler get_reminderByID_quary.Handler
}

func (x Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id := params.ByName("id")

	reminderID, err := reminder_id_model.NewReminderID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	reminder, err := x.queryHandler.Handle(r.Context(), get_reminderByID_quary.NewQuery(reminderID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	reminderForJson := postgresql_reminder_repository.NewReminderJson(reminder.GetID(),
		reminder.GetTitle(),
		reminder.GetDescription())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(reminderForJson)
}

func NewHandler(queryGetByIDHandler get_reminderByID_quary.Handler) Handler {
	return Handler{
		queryHandler: queryGetByIDHandler,
	}
}
