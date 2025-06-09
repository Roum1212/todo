package main

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	create_reminder_http_handler "github.com/Roum1212/todo/internal/api/http/handler/create-reminder"
	delete_reminder_http_handler "github.com/Roum1212/todo/internal/api/http/handler/delete-reminder"
	get_all_reminders_http_handler "github.com/Roum1212/todo/internal/api/http/handler/get-all-reminders"
	get_reminder_by_id_http_handler "github.com/Roum1212/todo/internal/api/http/handler/get-reminder-by-id"
	create_reminder_command "github.com/Roum1212/todo/internal/app/command/create-reminder"
	delete_reminder_command "github.com/Roum1212/todo/internal/app/command/delete-reminder"
	get_all_reminders_query "github.com/Roum1212/todo/internal/app/query/get-all-reminders"
	get_reminder_by_id_quary "github.com/Roum1212/todo/internal/app/query/get-reminder-by-id"
	postgresql_reminder_repository "github.com/Roum1212/todo/internal/infra/repository/reminder/postgresql"
	opentelemetry "github.com/Roum1212/todo/internal/pkg/opentelementry"
)

type Config struct {
	HTTPServer    ServerConfig `envPrefix:"HTTP_SERVER_"`
	OpenTelemetry ServerConfig `envPrefix:"OPENTELEMETRY_"`
	PostgreSQL    DBConfig     `envPrefix:"POSTGRESQL_"`
	Server        ServerConfig `envPrefix:"SERVER_"`
}

type DBConfig struct {
	DSN string `env:"DSN"`
}

type ServerConfig struct {
	Address string `env:"ADDRESS"`
	Version string `env:"VERSION"`
}

const (
	WriteTimeout = time.Second * 15
	ReadTimeout  = time.Second * 15
	IdleTimeout  = time.Second * 60
)

func main() {
	ctx := context.Background()

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		slog.ErrorContext(ctx, "failed to parse config", slog.Any("error", err))
		return //nolint:nlreturn // OK.
	}

	// OpenTelemetry | Resource.
	openTelemetryResource, err := opentelemetry.NewResource(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "failed to initialize OpenTelemetry resource", slog.Any("error", err))
		return //nolint:nlreturn // OK.
	}

	// OpenTelemetry | Logger Provider.
	openTelemetryLoggerProvider, err := opentelemetry.NewLoggerProvider(ctx, openTelemetryResource)
	if err != nil {
		slog.ErrorContext(ctx, "failed to initialize OpenTelemetry logger provider", slog.Any("error", err))
		return //nolint:nlreturn // OK.
	}

	defer func() {
		_ = openTelemetryLoggerProvider.Shutdown(ctx) //nolint:errcheck // OK.
	}()

	slog.SetDefault(otelslog.NewLogger(
		"reminder",
		otelslog.WithLoggerProvider(openTelemetryLoggerProvider),
		otelslog.WithSource(true),
	))

	// OpenTelemetry | Meter Provider.
	openTelemetryMeterProvider, err := opentelemetry.NewMeterProvider(ctx, openTelemetryResource)
	if err != nil {
		slog.ErrorContext(ctx, "failed to initialize OpenTelemetry meter", slog.Any("error", err))
		return //nolint:nlreturn // OK.
	}

	defer func() {
		_ = openTelemetryMeterProvider.Shutdown(ctx) //nolint:errcheck // OK.
	}()

	// OpenTelemetry | Tracer Provider.
	openTelemetryTracerProvider, err := opentelemetry.NewTracerProvider(ctx, openTelemetryResource)
	if err != nil {
		slog.ErrorContext(ctx, "failed to initialize OpenTelemetry tracer", slog.Any("error", err))
		return //nolint:nlreturn // OK.
	}

	defer func() {
		_ = openTelemetryTracerProvider.Shutdown(ctx) //nolint:errcheck // OK.
	}()

	// OpenTelemetry | Text Map Propagator.
	opentelemetry.SetTextMapPropagator()

	pool, err := pgxpool.New(ctx, cfg.PostgreSQL.DSN)
	if err != nil {
		slog.ErrorContext(ctx, "failed to initialize PostgreSQL connection", slog.Any("error", err))
		return //nolint:nlreturn // OK.
	}

	slog.InfoContext(ctx, "successfully initialized PostgreSQL connection")

	reminderRepository := postgresql_reminder_repository.NewRepository(pool)

	createReminderCommand := create_reminder_command.NewCommandHandler(reminderRepository)
	deleteReminderCommand := delete_reminder_command.NewCommandHandler(reminderRepository)
	getReminderByIDQuery := get_reminder_by_id_quary.NewQueryHandler(reminderRepository)
	getAllRemindersQuery := get_all_reminders_query.NewQueryHandler(reminderRepository)

	createReminderHTTPHandler := create_reminder_http_handler.NewHTTPHandler(createReminderCommand)
	deleteReminderHTTPHandler := delete_reminder_http_handler.NewHTTPHandler(deleteReminderCommand)
	getReminderByIDHTTPHandler := get_reminder_by_id_http_handler.NewHTTPHandler(getReminderByIDQuery)
	getAllRemindersHTTPHandler := get_all_reminders_http_handler.NewHTTPHandler(getAllRemindersQuery)

	router := httprouter.New()
	router.Handler(http.MethodPost, create_reminder_http_handler.Endpoint, createReminderHTTPHandler)
	router.Handler(http.MethodDelete, delete_reminder_http_handler.Endpoint, deleteReminderHTTPHandler)
	router.Handler(http.MethodGet, get_reminder_by_id_http_handler.Endpoint, getReminderByIDHTTPHandler)
	router.Handler(http.MethodGet, get_all_reminders_http_handler.Endpoint, getAllRemindersHTTPHandler)

	slog.InfoContext(
		ctx, "HTTP server starting",
		slog.String("address", cfg.HTTPServer.Address),
		slog.String("server version", cfg.Server.Version),
	)

	srv := &http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      otelhttp.NewHandler(cors.New(cors.Options{}).Handler(router), ""),
		WriteTimeout: WriteTimeout,
		ReadTimeout:  ReadTimeout,
		IdleTimeout:  IdleTimeout,
	}
	if err = srv.ListenAndServe(); err != nil {
		slog.Error("failed to start HTTP server", "error", err)
		return //nolint:nlreturn // OK.
	}
}
