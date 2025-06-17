package postgresql_reminder_repository

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/rueidis"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/proto"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	redis_reminder_repository "github.com/Roum1212/todo/internal/infra/repository/reminder/redis"
	reminder_v1 "github.com/Roum1212/todo/pkg/gen/reminder/v1"
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

func (x Repository) DeleteReminder(ctx context.Context, reminderID reminder_id_model.ReminderID) error {
	sql, args, err := squirrel.
		Delete(table).
		Where(squirrel.Eq{fieldID: int(reminderID)}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build sql: %w", err)
	}

	commandTag, err := x.client.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("failed to execute sql: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return reminder_aggregate.ErrReminderNotFound
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
		span.SetStatus(codes.Error, err.Error())

		return err
	}

	return nil
}

func (x tracerRepository) DeleteReminder(ctx context.Context, reminderID reminder_id_model.ReminderID) error {
	_, span := x.tracer.Start(ctx, "ReminderRepository.DeleteReminder")
	defer span.End()

	if err := x.repository.DeleteReminder(ctx, reminderID); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

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
		span.SetStatus(codes.Error, err.Error())

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

type redisRepository struct {
	repository reminder_aggregate.ReminderRepository
	client     rueidis.Client
}

func (x redisRepository) SaveReminder(ctx context.Context, reminder reminder_aggregate.Reminder) error {
	if err := x.repository.SaveReminder(ctx, reminder); err != nil {
		return err
	}

	if _, err := x.cacheAndGetReminder(ctx, reminder.GetID()); err != nil {
		return err
	}

	return nil
}

func (x redisRepository) DeleteReminder(ctx context.Context, reminderID reminder_id_model.ReminderID) error {
	if err := x.client.Do(
		ctx,
		x.client.B().Del().
			Key(redis_reminder_repository.NewKey(reminderID)).
			Build(),
	).Error(); err != nil {
		return fmt.Errorf("failed to delete reminder: %w", err)
	}

	if err := x.repository.DeleteReminder(ctx, reminderID); err != nil {
		return err
	}

	return nil
}

func (x redisRepository) GetAllReminders(ctx context.Context) ([]reminder_aggregate.Reminder, error) {
	return x.repository.GetAllReminders(ctx)
}

func (x redisRepository) GetReminderByID(
	ctx context.Context,
	reminderID reminder_id_model.ReminderID,
) (reminder_aggregate.Reminder, error) {
	v, err := x.client.Do(
		ctx,
		x.client.B().Get().
			Key(redis_reminder_repository.NewKey(reminderID)).
			Build(),
	).ToString()
	if err != nil {
		switch {
		case rueidis.IsRedisNil(err):
			return x.cacheAndGetReminder(ctx, reminderID)
		default:
			return reminder_aggregate.Reminder{}, fmt.Errorf("failed to get reminder: %w", err)
		}
	}

	decode, err := base64.StdEncoding.DecodeString(v)
	if err != nil {
		return reminder_aggregate.Reminder{}, fmt.Errorf("failed to decode reminder: %w", err)
	}

	var reminderDTO reminder_v1.Reminder
	if err = proto.Unmarshal(decode, &reminderDTO); err != nil {
		return reminder_aggregate.Reminder{}, fmt.Errorf("failed to unmarshal reminder: %w", err)
	}

	reminder, err := redis_reminder_repository.ToReminder(&reminderDTO)
	if err != nil {
		return reminder_aggregate.Reminder{}, fmt.Errorf("failed to create reminder: %w", err)
	}

	return reminder, nil
}

func (x redisRepository) cacheAndGetReminder(
	ctx context.Context,
	reminderID reminder_id_model.ReminderID,
) (reminder_aggregate.Reminder, error) {
	reminder, err := x.repository.GetReminderByID(ctx, reminderID)
	if err != nil {
		return reminder_aggregate.Reminder{}, err
	}

	reminderDTO := redis_reminder_repository.NewReminderDTO(reminder)

	data, err := proto.Marshal(reminderDTO)
	if err != nil {
		return reminder_aggregate.Reminder{}, fmt.Errorf("failed to marshal proto: %w", err)
	}

	encode := base64.StdEncoding.EncodeToString(data)
	if err = x.client.Do(
		ctx,
		x.client.B().Set().
			Key(redis_reminder_repository.NewKey(reminderID)).
			Value(encode).
			Build(),
	).Error(); err != nil {
		return reminder_aggregate.Reminder{}, fmt.Errorf("failed to save reminder: %w", err)
	}

	return reminder, nil
}

func NewRepositoryWithRedis(
	repository reminder_aggregate.ReminderRepository,
	client rueidis.Client,
) reminder_aggregate.ReminderRepository {
	return redisRepository{
		repository: repository,
		client:     client,
	}
}
