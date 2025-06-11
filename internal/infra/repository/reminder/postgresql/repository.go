package postgresql_reminder_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

const table = "reminders"

const (
	fieldID          = "id"
	fieldTitle       = "title"
	fieldDescription = "description"
)

const tracerName = "github.com/Roum1212/todo/internal/postgresql/reminder/repository"

type Repository struct {
	client *pgxpool.Pool
}

func (x Repository) SaveReminder(ctx context.Context, reminder reminder_aggregate.Reminder) error {
	reminderDTO := NewReminder(reminder)

	sql, args, err := squirrel.
		Insert(table).
		Columns(fieldID, fieldTitle, fieldDescription).
		Values(reminderDTO.ID, reminderDTO.Title, reminderDTO.Description).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build sql: %w", err)
	}

	if _, err = x.client.Exec(ctx, sql, args...); err != nil {
		return fmt.Errorf("failed to execute sql: %w", err)
	}

	return nil
}

func (x Repository) DeleteReminder(
	ctx context.Context,
	reminderID reminder_id_model.ReminderID,
) error {
	sql, args, err := squirrel.
		Delete(table).
		Where(squirrel.Eq{fieldID: int(reminderID)}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build sql: %w", err)
	}

	if _, err = x.client.Exec(ctx, sql, args...); err != nil {
		return fmt.Errorf("failed to execute sql: %w", err)
	}

	return nil
}

func (x Repository) GetReminderByID(
	ctx context.Context,
	reminderID reminder_id_model.ReminderID,
) (reminder_aggregate.Reminder, error) {
	var reminderDTO Reminder

	sql, args, err := squirrel.
		Select(fieldID, fieldTitle, fieldDescription).
		From(table).
		Where(squirrel.Eq{fieldID: reminderID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return reminder_aggregate.Reminder{}, fmt.Errorf("failed to get sql: %w", err)
	}

	if err = pgxscan.Get(ctx, x.client, &reminderDTO, sql, args...); err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return reminder_aggregate.Reminder{}, reminder_aggregate.ErrReminderNotFound
		default:
			return reminder_aggregate.Reminder{}, fmt.Errorf("failed to get reminder: %w", err)
		}
	}

	reminder, err := ToReminder(reminderDTO)
	if err != nil {
		return reminder_aggregate.Reminder{}, fmt.Errorf("failed to create reminder: %w", err)
	}

	return reminder, nil
}

func (x Repository) GetAllReminders(ctx context.Context) ([]reminder_aggregate.Reminder, error) {
	var reminderDTOs []Reminder

	sql, args, err := squirrel.
		Select(fieldID, fieldTitle, fieldDescription).
		From(table).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build sql: %w", err)
	}

	if err = pgxscan.Select(ctx, x.client, &reminderDTOs, sql, args...); err != nil {
		return nil, fmt.Errorf("failed to query sql: %w", err)
	}

	if len(reminderDTOs) == 0 {
		return nil, reminder_aggregate.ErrReminderNotFound
	}

	reminders, err := ToReminders(reminderDTOs...)
	if err != nil {
		return nil, fmt.Errorf("failed to create reminders: %w", err)
	}

	return reminders, nil
}

func NewRepository(client *pgxpool.Pool) reminder_aggregate.ReminderRepository {
	return Repository{
		client: client,
	}
}

type tracerRepository struct {
	repository reminder_aggregate.ReminderRepository
	tracer     trace.Tracer
}

func (x tracerRepository) SaveReminder(ctx context.Context, reminder reminder_aggregate.Reminder) error {
	_, span := x.tracer.Start(ctx, "ReminderRepository.SaveReminder")
	defer span.End()

	if err := x.repository.SaveReminder(ctx, reminder); err != nil {
		span.RecordError(err)

		return err
	}

	return nil
}

func (x tracerRepository) DeleteReminder(ctx context.Context, reminderID reminder_id_model.ReminderID) error {
	_, span := x.tracer.Start(ctx, "ReminderRepository.DeleteReminder")
	defer span.End()

	if err := x.repository.DeleteReminder(ctx, reminderID); err != nil {
		span.RecordError(err)

		return err
	}

	return nil
}

func (x tracerRepository) GetAllReminders(ctx context.Context) ([]reminder_aggregate.Reminder, error) {
	_, span := x.tracer.Start(ctx, "ReminderRepository.GetAllReminders")
	defer span.End()

	reminders, err := x.repository.GetAllReminders(ctx)
	if err != nil {
		span.RecordError(err)

		return nil, err
	}

	return reminders, nil
}

func (x tracerRepository) GetReminderByID(
	ctx context.Context,
	reminderID reminder_id_model.ReminderID,
) (reminder_aggregate.Reminder, error) {
	_, span := x.tracer.Start(ctx, "ReminderRepository.GetReminderByID")
	defer span.End()

	reminder, err := x.repository.GetReminderByID(ctx, reminderID)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return reminder_aggregate.Reminder{}, err
	}

	return reminder, nil
}

func NewRepositoryWithTracing(repository reminder_aggregate.ReminderRepository) reminder_aggregate.ReminderRepository {
	return tracerRepository{
		repository: repository,
		tracer:     otel.Tracer(tracerName),
	}
}
