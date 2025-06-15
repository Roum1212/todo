package main

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	create_reminder_rpc "github.com/Roum1212/todo/internal/api/grpc/rpc/create-reminder"
	delete_reminder_rpc "github.com/Roum1212/todo/internal/api/grpc/rpc/delete-reminder"
	get_all_reminders_rpc "github.com/Roum1212/todo/internal/api/grpc/rpc/get-all-reminders"
	get_reminder_by_id_rpc "github.com/Roum1212/todo/internal/api/grpc/rpc/get-reminder-by-id"
	grpc_server "github.com/Roum1212/todo/internal/api/grpc/server"
	create_reminder_http_handler "github.com/Roum1212/todo/internal/api/http/handler/create-reminder"
	delete_reminder_http_handler "github.com/Roum1212/todo/internal/api/http/handler/delete-reminder"
	get_all_reminders_http_handler "github.com/Roum1212/todo/internal/api/http/handler/get-all-reminders"
	get_reminder_by_id_http_handler "github.com/Roum1212/todo/internal/api/http/handler/get-reminder-by-id"
	create_reminder_command "github.com/Roum1212/todo/internal/app/command/create-reminder"
	delete_reminder_command "github.com/Roum1212/todo/internal/app/command/delete-reminder"
	get_all_reminders_query "github.com/Roum1212/todo/internal/app/query/get-all-reminders"
	get_reminder_by_id_query "github.com/Roum1212/todo/internal/app/query/get-reminder-by-id"
	postgresql_reminder_repository "github.com/Roum1212/todo/internal/infra/repository/reminder/postgresql"
	opentelemetry "github.com/Roum1212/todo/internal/pkg/opentelementry"
	reminder_v1 "github.com/Roum1212/todo/pkg/gen/reminder/v1"
)

type Config struct {
	GRPCServer    ServerConfig `envPrefix:"GRPC_SERVER_"`
	HTTPServer    ServerConfig `envPrefix:"HTTP_SERVER_"`
	OpenTelemetry ServerConfig `envPrefix:"OPENTELEMETRY_"`
	PostgreSQL    DBConfig     `envPrefix:"POSTGRESQL_"`
}

type DBConfig struct {
	DSN string `env:"DSN"`
}

type ServerConfig struct {
	Address string `env:"ADDRESS"`
}

const (
	WriteTimeout = time.Second * 15
	ReadTimeout  = time.Second * 15
	IdleTimeout  = time.Second * 60
)

func main() { //nolint:gocognit,cyclop // OK.
	ctx := context.Background()

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		slog.ErrorContext(ctx, "failed to parse config", slog.Any("error", err))

		return
	}

	// OpenTelemetry | Resource.
	openTelemetryResource, err := opentelemetry.NewResource(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "failed to create opentelemetry resource", slog.Any("error", err))

		return
	}

	// OpenTelemetry | Logger Provider.
	openTelemetryLoggerProvider, err := opentelemetry.NewLoggerProvider(ctx, openTelemetryResource)
	if err != nil {
		slog.ErrorContext(ctx, "failed to create opentelemetry logger provider", slog.Any("error", err))

		return
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
		slog.ErrorContext(ctx, "failed to create opentelemetry meter provider", slog.Any("error", err))

		return
	}

	defer func() {
		_ = openTelemetryMeterProvider.Shutdown(ctx) //nolint:errcheck // OK.
	}()

	// OpenTelemetry | Tracer Provider.
	openTelemetryTracerProvider, err := opentelemetry.NewTracerProvider(ctx, openTelemetryResource)
	if err != nil {
		slog.ErrorContext(ctx, "failed to create opentelemetry tracer provider", slog.Any("error", err))

		return
	}

	defer func() {
		_ = openTelemetryTracerProvider.Shutdown(ctx) //nolint:errcheck // OK.
	}()

	// OpenTelemetry | Text Map Propagator.
	opentelemetry.SetTextMapPropagator()

	pool, err := pgxpool.New(ctx, cfg.PostgreSQL.DSN)
	if err != nil {
		slog.ErrorContext(ctx, "failed to create postgresql pool", slog.Any("error", err))

		return
	}

	reminderRepository := postgresql_reminder_repository.NewRepository(pool)
	reminderRepository = postgresql_reminder_repository.NewRepositoryWithTracing(reminderRepository)

	createReminderCommand := create_reminder_command.NewCommandHandler(reminderRepository)
	createReminderCommand = create_reminder_command.NewCommandHandlerWithTracing(createReminderCommand)

	deleteReminderCommand := delete_reminder_command.NewCommandHandler(reminderRepository)
	deleteReminderCommand = delete_reminder_command.NewCommandHandlerTracer(deleteReminderCommand)

	getAllRemindersQuery := get_all_reminders_query.NewQueryHandler(reminderRepository)
	getAllRemindersQuery = get_all_reminders_query.NewQueryHandlerTracer(getAllRemindersQuery)

	getReminderByIDQuery := get_reminder_by_id_query.NewQueryHandler(reminderRepository)
	getReminderByIDQuery = get_reminder_by_id_query.NewQueryHandlerTracer(getReminderByIDQuery)

	createReminderHTTPHandler := create_reminder_http_handler.NewHTTPHandler(createReminderCommand)
	deleteReminderHTTPHandler := delete_reminder_http_handler.NewHTTPHandler(deleteReminderCommand)
	getAllRemindersHTTPHandler := get_all_reminders_http_handler.NewHTTPHandler(getAllRemindersQuery)
	getReminderByIDHTTPHandler := get_reminder_by_id_http_handler.NewHTTPHandler(getReminderByIDQuery)

	router := httprouter.New()
	router.Handler(http.MethodPost, create_reminder_http_handler.Endpoint, createReminderHTTPHandler)
	router.Handler(http.MethodDelete, delete_reminder_http_handler.Endpoint, deleteReminderHTTPHandler)
	router.Handler(http.MethodGet, get_reminder_by_id_http_handler.Endpoint, getReminderByIDHTTPHandler)
	router.Handler(http.MethodGet, get_all_reminders_http_handler.Endpoint, getAllRemindersHTTPHandler)

	srv := &http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      otelhttp.NewHandler(cors.New(cors.Options{}).Handler(router), "HTTP Server"),
		WriteTimeout: WriteTimeout,
		ReadTimeout:  ReadTimeout,
		IdleTimeout:  IdleTimeout,
	}

	listen, err := net.Listen("tcp", cfg.GRPCServer.Address)
	if err != nil {
		slog.ErrorContext(ctx, "failed to listen grpc server", slog.Any("error", err))

		return
	}

	grpcServer := grpc.NewServer(grpc.StatsHandler(otelgrpc.NewServerHandler(
		otelgrpc.WithSpanOptions(
			trace.WithSpanKind(trace.SpanKindServer),
			trace.WithNewRoot(),
		),
	)))

	reminder_v1.RegisterReminderServiceServer(
		grpcServer,
		grpc_server.NewCreateReminderService(
			create_reminder_rpc.NewCreateReminderRPC(createReminderCommand),
			delete_reminder_rpc.NewDeleteReminderRPC(deleteReminderCommand),
			get_all_reminders_rpc.NewGetAllRemindersRPC(getAllRemindersQuery),
			get_reminder_by_id_rpc.NewGetReminderByIDRPC(getReminderByIDQuery),
		),
	)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		if err = srv.ListenAndServe(); err != nil {
			slog.Error("failed to listen and serve http server", slog.Any("error", err))

			return err
		}

		return nil
	})

	g.Go(func() error {
		if err = grpcServer.Serve(listen); err != nil {
			slog.ErrorContext(ctx, "failed to serve grpc server", slog.Any("error", err))

			return err
		}

		return nil
	})

	if err = g.Wait(); err != nil {
		slog.ErrorContext(ctx, "failed to shutdown server", slog.Any("error", err))

		return
	}
}
