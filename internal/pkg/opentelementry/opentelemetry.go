package opentelemetry

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

func NewLoggerProvider(
	ctx context.Context,
	res *resource.Resource,
) (*log.LoggerProvider, error) {
	exporter, err := otlploggrpc.New(
		ctx,
		otlploggrpc.WithEndpoint("otel-collector:4317"),
		otlploggrpc.WithInsecure(),
		otlploggrpc.WithRetry(otlploggrpc.RetryConfig{
			Enabled:         true,
			InitialInterval: time.Second * 5,
			MaxElapsedTime:  time.Minute,
			MaxInterval:     time.Second * 30,
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("otlploggrpc.New: %w", err)
	}

	x := log.NewLoggerProvider(
		log.WithProcessor(log.NewBatchProcessor(exporter, log.WithExportBufferSize(1))),
		log.WithResource(res),
	)

	global.SetLoggerProvider(x)

	return x, nil
}

func NewMeterProvider(
	ctx context.Context,
	res *resource.Resource,
) (*metric.MeterProvider, error) {
	exporter, err := otlpmetricgrpc.New(
		ctx,
		otlpmetricgrpc.WithEndpoint("otel-collector:4317"),
		otlpmetricgrpc.WithInsecure(),
		// otlpmetricgrpc .WithHeaders(),
		otlpmetricgrpc.WithRetry(otlpmetricgrpc.RetryConfig{
			Enabled:         true,
			InitialInterval: time.Second * 5,
			MaxElapsedTime:  time.Minute,
			MaxInterval:     time.Second * 30,
		}),
		// otlpmetricgrpc .WithTemporalitySelector(),
	)
	if err != nil {
		return nil, fmt.Errorf("otlpmetricgrpc.New: %w", err)
	}

	x := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(exporter)),
		metric.WithResource(res),
		// metric.WithView(),
	)

	otel.SetMeterProvider(x)

	return x, nil
}

func NewTracerProvider(
	ctx context.Context,
	res *resource.Resource,
) (*trace.TracerProvider, error) {
	exporter, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithEndpoint("otel-collector:4317"),
		otlptracegrpc.WithInsecure(),
		// otlptracegrpc .WithHeaders(),
		otlptracegrpc.WithRetry(otlptracegrpc.RetryConfig{
			Enabled:         true,
			InitialInterval: time.Second * 5,
			MaxElapsedTime:  time.Minute,
			MaxInterval:     time.Second * 30,
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("otlptracegrpc.New: %w", err)
	}

	x := trace.NewTracerProvider(
		trace.WithBatcher(
			exporter,
			trace.WithBatchTimeout(time.Second*5),
			// trace.WithBlocking(),
			trace.WithExportTimeout(time.Second*30),
			trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
			trace.WithMaxQueueSize(trace.DefaultMaxQueueSize),
		),
		// trace.WithIDGenerator(), // TODO: UUID v7.
		trace.WithResource(res),
		trace.WithSampler(trace.AlwaysSample()),
	)

	otel.SetTracerProvider(x)

	return x, nil
}

func NewResource(ctx context.Context) (*resource.Resource, error) {
	return resource.New(
		ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String("reminder"),
			semconv.ServiceVersionKey.String("1.0.0"),
			semconv.DeploymentEnvironmentKey.String("development"),
		),
		resource.WithContainer(),
		resource.WithFromEnv(),
		resource.WithHost(),
		// info.WithHostID(),
		resource.WithOS(),
		// info.WithSchemaURL(semconv.SchemaURL),
		resource.WithTelemetrySDK(),
	)
}

func SetTextMapPropagator() {
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.Baggage{},
		propagation.TraceContext{},
	))
}
