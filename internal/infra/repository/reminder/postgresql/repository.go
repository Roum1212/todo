package postgresql_reminder_repository

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"

	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
)

const table = "reminders"

const (
	fieldID          = "id"
	fieldTitle       = "title"
	fieldDescription = "description"
)

type Repository struct {
	client *pgxpool.Pool
}

func (x Repository) SaveReminder(ctx context.Context, reminder reminder_aggregate.Reminder) error {
	sql, args, err := squirrel.
		Insert(table).
		Columns(fieldID, fieldTitle, fieldDescription).
		Values(int(reminder.GetID()), string(reminder.GetTitle()), string(reminder.GetDescription())).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("faild to build sql: %w", err)
	}

	if _, err = x.client.Exec(ctx, sql, args...); err != nil {
		return fmt.Errorf("faild to execute sql: %w", err)
	}

	return nil
}

func (x Repository) DeleteReminder(
	ctx context.Context,
	reminderID reminder_id_model.ReminderID,
) error {
	log.Println("The reminder with the ID", reminderID, "has been deleted")

	return nil
}

func (x Repository) GetReminderByID(
	ctx context.Context,
	reminderID reminder_id_model.ReminderID,
) (reminder_aggregate.Reminder, error) {
	log.Println("The reminder with the ID", reminderID, "has been found")

	reminder := reminder_aggregate.Reminder{}
	return reminder, nil
}

func (x Repository) GetAllReminders(ctx context.Context) ([]reminder_aggregate.Reminder, error) {
	var remindersDTOs []Reminder
	var ErrRemindersNotFound = errors.New("reminders are not found")

	sql, args, err := squirrel.
		Select(fieldID, fieldTitle, fieldDescription).
		From(table).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("faild to build sql: %w", err)
	}

	if err = pgxscan.Select(ctx, x.client, &remindersDTOs, sql, args...); err != nil {
		return nil, fmt.Errorf("faild to query sql: %w", err)
	}

	if len(remindersDTOs) == 0 {
		return nil, ErrRemindersNotFound
	}

	reminders := make([]reminder_aggregate.Reminder, len(remindersDTOs))
	for i, reminderDTO := range remindersDTOs {
		reminders[i] = reminder_aggregate.NewReminder(
			reminder_id_model.ReminderID(reminderDTO.ID),
			reminder_title_model.NewReminderTitle(reminderDTO.Title),
			reminder_description_model.NewReminderDescription(reminderDTO.Description),
		)

	}

	return reminders, nil
}

func NewRepository(client *pgxpool.Pool) Repository {
	return Repository{
		client: client,
	}
}
