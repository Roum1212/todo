package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/caarlos0/env/v11"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"

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

type Config struct {
	PostgreSQLDSN string `env:"POSTGRESQL_DSN"`
	DBNAME        string `env:"POSTGRESQL_DBNAME"`
}

func main() {
	ctx := context.Background()

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal("failed to parse env: %w", err)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("failed to set goose dialect: %v", err)
	}

	postgres, err := sql.Open("postgres", cfg.PostgreSQLDSN)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	defer postgres.Close()

	if err = postgres.Ping(); err != nil {
		log.Fatalf("failed to ping postgres: %v", err)
	}

	if err = goose.Up(postgres, "/app/migrations"); err != nil {
		log.Fatalf("failed to up migrations: %v", err)
	}

	pool, err := pgxpool.New(ctx, cfg.PostgreSQLDSN)
	if err != nil {
		log.Fatal("failed to create pgx pool: %w", err)
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
