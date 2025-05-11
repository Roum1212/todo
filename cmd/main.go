package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	create_reminder_http_handler "github.com/Roum1212/todo/internal/api/http/handler/create-reminder"
	create_reminder_command "github.com/Roum1212/todo/internal/app/command/create-reminder"
	postgresql_reminder_repository "github.com/Roum1212/todo/internal/infra/repository/reminder/postgresql"
)

func main() {
	reminderRepository := postgresql_reminder_repository.NewRepository()
	createReminderCommand := create_reminder_command.NewHandler(reminderRepository)
	createReminderHTTPHandler := create_reminder_http_handler.NewHandler(createReminderCommand)

	router := httprouter.New()
	router.Handler(http.MethodPost, create_reminder_http_handler.Endpoint, createReminderHTTPHandler)

	http.ListenAndServe(":9080", router)
}
