package create_reminder_rpc

import (
	"context"
	"crypto/rand"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	create_reminder_command "github.com/Roum1212/todo/internal/app/command/create-reminder"
	create_reminder_command_mock "github.com/Roum1212/todo/internal/app/command/create-reminder/mock"
	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
	reminder_v1 "github.com/Roum1212/todo/pkg/gen/reminder/v1"
)

func TestCreateReminder(t *testing.T) {
	t.Parallel()

	request := reminder_v1.CreateReminderRequest{
		Title:       rand.Text(),
		Description: rand.Text(),
	}

	mc := minimock.NewController(t)

	title, err := reminder_title_model.NewReminderTitle(request.Title)
	require.NoError(t, err)

	description, err := reminder_description_model.NewReminderDescription(request.Description)
	require.NoError(t, err)

	commandHandlerMock := create_reminder_command_mock.NewCommandHandlerMock(mc).
		HandleCommandMock.
		Inspect(func(ctx context.Context, c create_reminder_command.Command) {
			require.Equal(t, title, c.GetTitle())
			require.Equal(t, description, c.GetDescription())
		}).
		Return(nil)

	createReminderRPC := NewCreateReminderRPC(commandHandlerMock)

	createReminderResponse, err := createReminderRPC.CreateReminder(t.Context(), &request)
	require.NoError(t, err)
	require.NotNil(t, createReminderResponse)
}

func TestCreateReminder_InvalidArgument(t *testing.T) {
	t.Parallel()

	t.Run("invalid title", func(t *testing.T) {
		t.Parallel()

		request := reminder_v1.CreateReminderRequest{
			Title:       "",
			Description: rand.Text(),
		}

		createReminderRPC := NewCreateReminderRPC(nil)

		createReminderResponse, err := createReminderRPC.CreateReminder(t.Context(), &request)
		require.Error(t, err)
		require.Nil(t, createReminderResponse)
	})

	t.Run("invalid description", func(t *testing.T) {
		t.Parallel()

		request := reminder_v1.CreateReminderRequest{
			Title:       rand.Text(),
			Description: "",
		}

		createReminderRPC := NewCreateReminderRPC(nil)

		createReminderResponse, err := createReminderRPC.CreateReminder(t.Context(), &request)
		require.Error(t, err)
		require.Nil(t, createReminderResponse)
	})
}

func TestCreateReminder_Internal(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	request := reminder_v1.CreateReminderRequest{
		Title:       rand.Text(),
		Description: rand.Text(),
	}

	title, err := reminder_title_model.NewReminderTitle(request.Title)
	require.NoError(t, err)

	description, err := reminder_description_model.NewReminderDescription(request.Description)
	require.NoError(t, err)

	commandHandlerMock := create_reminder_command_mock.NewCommandHandlerMock(mc).
		HandleCommandMock.
		Inspect(func(ctx context.Context, c create_reminder_command.Command) {
			require.Equal(t, title, c.GetTitle())
			require.Equal(t, description, c.GetDescription())
		}).
		Return(status.Errorf(codes.Internal, "internal error: %v", err))

	createReminderRPC := NewCreateReminderRPC(commandHandlerMock)

	createReminderResponse, err := createReminderRPC.CreateReminder(t.Context(), &request)
	require.Nil(t, createReminderResponse)

	st, _ := status.FromError(err)
	require.Equal(t, codes.Internal, st.Code())
}
