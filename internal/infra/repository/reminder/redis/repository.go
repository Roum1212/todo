package redis_reminder_repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/redis/rueidis"
	"google.golang.org/protobuf/proto"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	reminder_v1 "github.com/Roum1212/todo/pkg/gen/reminder/v1"
)

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
	if err := x.repository.DeleteReminder(ctx, reminderID); err != nil {
		return err
	}

	if err := x.client.Do(
		ctx,
		x.client.B().Del().
			Key(buildReminderKey(reminderID)).
			Build(),
	).Error(); err != nil {
		return fmt.Errorf("failed to delete reminder: %w", err)
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
			Key(buildReminderKey(reminderID)).
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

	var reminderDTO reminder_v1.Reminder
	if err = proto.Unmarshal([]byte(v), &reminderDTO); err != nil {
		return reminder_aggregate.Reminder{}, fmt.Errorf("failed to unmarshal reminder: %w", err)
	}

	reminder, err := ToReminder(&reminderDTO)
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

	reminderDTO := NewReminderDTO(reminder)

	data, err := proto.Marshal(reminderDTO)
	if err != nil {
		return reminder_aggregate.Reminder{}, fmt.Errorf("failed to marshal proto: %w", err)
	}

	if err = x.client.Do(
		ctx,
		x.client.B().Set().
			Key(buildReminderKey(reminderID)).
			Value(string(data)).
			Build(),
	).Error(); err != nil {
		return reminder_aggregate.Reminder{}, fmt.Errorf("failed to save reminder: %w", err)
	}

	return reminder, nil
}

func buildReminderKey(reminderID reminder_id_model.ReminderID) string {
	return strconv.FormatInt(int64(reminderID), 10)
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
