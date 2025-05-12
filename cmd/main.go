package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	create_reminder_http_handler "github.com/Roum1212/todo/internal/api/http/handler/create-reminder"
	delete_reminder_http_handler "github.com/Roum1212/todo/internal/api/http/handler/delete-reminder"
	create_reminder_command "github.com/Roum1212/todo/internal/app/command/create-reminder"
	"github.com/Roum1212/todo/internal/app/command/delete-reminder"

	postgresql_reminder_repository "github.com/Roum1212/todo/internal/infra/repository/reminder/postgresql"
)

func main() {
	reminderRepository := postgresql_reminder_repository.NewRepository()

	createReminderCommand := create_reminder_command.NewHandler(reminderRepository)
	createReminderHTTPHandler := create_reminder_http_handler.NewHandler(createReminderCommand)

	deleteReminderCommand := delete_reminder_command.DeleteHandler(reminderRepository)
	deleteReminderHTTPHandler := delete_reminder_http_handler.NewHandler(deleteReminderCommand)

	router := httprouter.New()
	router.Handler(http.MethodPost, create_reminder_http_handler.Endpoint, createReminderHTTPHandler)
	router.DELETE(
		delete_reminder_http_handler.Endpoint,
		func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			deleteReminderHTTPHandler.ServeHTTP(w, r, ps)
		},
	)

	if err := http.ListenAndServe(":9080", router); err != nil {
		log.Fatal(err)
	}
}
