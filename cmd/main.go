package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	create_reminder_http_handler "github.com/Roum1212/todo/internal/api/http/handler/create-reminder"
	delete_reminder_http_handler "github.com/Roum1212/todo/internal/api/http/handler/delete-reminder"
	get_all_reminders_http_handler "github.com/Roum1212/todo/internal/api/http/handler/get-all-reminders"
	get_reminderByID_http_handler "github.com/Roum1212/todo/internal/api/http/handler/get-reminder-by-id"
	create_reminder_command "github.com/Roum1212/todo/internal/app/command/create-reminder"
	delete_reminder_command "github.com/Roum1212/todo/internal/app/command/delete-reminder"
	get_all_reminders_quer "github.com/Roum1212/todo/internal/app/query/get-all-reminders"
	get_reminderByID_quary "github.com/Roum1212/todo/internal/app/query/get-reminder-by-id"
	postgresql_reminder_repository "github.com/Roum1212/todo/internal/infra/repository/reminder/postgresql"
)

func main() {
	reminderRepository := postgresql_reminder_repository.NewRepository()

	createReminderCommand := create_reminder_command.NewHandler(reminderRepository)
	deleteReminderCommand := delete_reminder_command.NewHandler(reminderRepository)
	getReminderByIdQuery := get_reminderByID_quary.NewHandler(reminderRepository)
	getAllReminders := get_all_reminders_quer.NewHandler(reminderRepository)

	createReminderHTTPHandler := create_reminder_http_handler.NewHandler(createReminderCommand)
	deleteReminderHTTPHandler := delete_reminder_http_handler.NewHandler(deleteReminderCommand)
	getReminderByIdHTTPHandler := get_reminderByID_http_handler.NewHandler(getReminderByIdQuery)
	getAllRemindersHTTPHandler := get_all_reminders_http_handler.NewHandler(getAllReminders)

	router := httprouter.New()
	router.Handler(http.MethodPost, create_reminder_http_handler.Endpoint, createReminderHTTPHandler)
	router.Handler(http.MethodDelete, delete_reminder_http_handler.Endpoint, deleteReminderHTTPHandler)
	router.Handler(http.MethodGet, get_reminderByID_http_handler.Endpoint, getReminderByIdHTTPHandler)
	router.Handler(http.MethodGet, get_all_reminders_http_handler.Endpoint, getAllRemindersHTTPHandler)

	if err := http.ListenAndServe(":9080", router); err != nil {
		log.Fatal(err)
	}
}
