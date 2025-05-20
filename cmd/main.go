package main

import (
	"context"
	"log"
	"net/http"

	"github.com/caarlos0/env/v11"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"

	create_reminder_http_handler "github.com/Roum1212/todo/internal/api/http/handler/create-reminder"
	delete_reminder_http_handler "github.com/Roum1212/todo/internal/api/http/handler/delete-reminder"
	get_all_reminders_http_handler "github.com/Roum1212/todo/internal/api/http/handler/get-all-reminders"
	get_reminder_by_id_http_handler "github.com/Roum1212/todo/internal/api/http/handler/get-reminder-by-id"
	create_reminder_command "github.com/Roum1212/todo/internal/app/command/create-reminder"
	delete_reminder_command "github.com/Roum1212/todo/internal/app/command/delete-reminder"
	get_all_reminders_query "github.com/Roum1212/todo/internal/app/query/get-all-reminders"
	get_reminder_by_id_quary "github.com/Roum1212/todo/internal/app/query/get-reminder-by-id"
	postgresql_reminder_repository "github.com/Roum1212/todo/internal/infra/repository/reminder/postgresql"
)

type Сonfig struct {
	PostgresDSN string `env:"POSTGRESQL_DSN"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := Сonfig{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal("failed to parse env")
	}

	pool, err := pgxpool.New(context.Background(), cfg.PostgresDSN)
	if err != nil {
		log.Fatal(err)
	}

	reminderRepository := postgresql_reminder_repository.NewRepository(pool)

	createReminderCommand := create_reminder_command.NewHandler(reminderRepository)
	deleteReminderCommand := delete_reminder_command.NewHandler(reminderRepository)
	getReminderByIDQuery := get_reminder_by_id_quary.NewHandler(reminderRepository)
	getAllRemindersQuery := get_all_reminders_query.NewHandler(reminderRepository)

	createReminderHTTPHandler := create_reminder_http_handler.NewHandler(createReminderCommand)
	deleteReminderHTTPHandler := delete_reminder_http_handler.NewHandler(deleteReminderCommand)
	getReminderByIDHTTPHandler := get_reminder_by_id_http_handler.NewHandler(getReminderByIDQuery)
	getAllRemindersHTTPHandler := get_all_reminders_http_handler.NewHandler(getAllRemindersQuery)

	router := httprouter.New()
	router.Handler(http.MethodPost, create_reminder_http_handler.Endpoint, createReminderHTTPHandler)
	router.Handler(http.MethodDelete, delete_reminder_http_handler.Endpoint, deleteReminderHTTPHandler)
	router.Handler(http.MethodGet, get_reminder_by_id_http_handler.Endpoint, getReminderByIDHTTPHandler)
	router.Handler(http.MethodGet, get_all_reminders_http_handler.Endpoint, getAllRemindersHTTPHandler)

	if err = http.ListenAndServe(":9080", router); err != nil {
		log.Fatal(err)
	}
}
