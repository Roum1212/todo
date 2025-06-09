package get_reminder_by_id_http_handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"

	get_reminder_by_id_quary "github.com/Roum1212/todo/internal/app/query/get-reminder-by-id"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

const Endpoint = "/reminders/" + paramID

const paramID = ":id"

type Handler struct {
	queryHandler get_reminder_by_id_quary.QueryHandler
}

func (x Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slog.InfoContext(r.Context(), "handle request", slog.Any("request", r))
	params := httprouter.ParamsFromContext(r.Context())

	reminderID, err := reminder_id_model.NewReminderIDFromString(
		params.ByName(strings.TrimPrefix(paramID, ":")),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		slog.WarnContext(
			r.Context(),
			"incorrect reminder id",
			slog.Any("error", err),
			slog.String("URL", r.URL.String()),
			slog.String("method", r.Method),
		)

		return
	}

	reminder, err := x.queryHandler.HandleQuery(
		r.Context(),
		get_reminder_by_id_quary.NewQuery(reminderID),
	)
	if err != nil {
		switch {
		case errors.Is(err, get_reminder_by_id_quary.ErrReminderNotFound):
			http.Error(w, err.Error(), http.StatusNotFound)
			slog.WarnContext(
				r.Context(),
				"failed to find reminder",
				slog.Any("error", err),
				slog.String("URL", r.URL.String()),
				slog.String("Method", r.Method),
			)

			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			slog.ErrorContext(
				r.Context(),
				"failed to process request",
				slog.Any("error", err),
				slog.String("URL", r.URL.String()),
				slog.String("Method", r.Method),
			)

			return
		}
	}

	reminderDTO := NewReminder(reminder)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(reminderDTO) //nolint:errcheck,errchkjson // OK.
	{
		slog.InfoContext(
			r.Context(),
			"successfully retrieved reminder",
			slog.String("URL", r.URL.String()),
			slog.String("method", r.Method),
		)
	}
}

func NewHTTPHandler(queryHandler get_reminder_by_id_quary.QueryHandler) Handler {
	return Handler{
		queryHandler: queryHandler,
	}
}
