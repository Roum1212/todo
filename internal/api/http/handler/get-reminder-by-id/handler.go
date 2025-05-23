package get_reminder_by_id_http_handler

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	get_reminder_by_id_quary "github.com/Roum1212/todo/internal/app/query/get-reminder-by-id"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

const Endpoint = "/reminders/:id"

type Handler struct {
	queryHandler get_reminder_by_id_quary.Handler
}

func (x Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id := params.ByName("id")

	reminderID, err := reminder_id_model.NewReminderID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	reminder, err := x.queryHandler.Handle(r.Context(), get_reminder_by_id_quary.NewQuery(reminderID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	reminderDTO := NewReminder(
		reminder.GetID(),
		reminder.GetTitle(),
		reminder.GetDescription(),
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(reminderDTO) //nolint:errcheck,errchkjson // OK.
}

func NewHandler(queryHandler get_reminder_by_id_quary.Handler) Handler {
	return Handler{
		queryHandler: queryHandler,
	}
}
