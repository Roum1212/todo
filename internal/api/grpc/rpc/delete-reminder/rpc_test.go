package delete_reminder_rpc

import (
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	delete_reminder_command "github.com/Roum1212/todo/internal/app/command/delete-reminder"
	delete_reminder_command_mock "github.com/Roum1212/todo/internal/app/command/delete-reminder/mock"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	reminder_v1 "github.com/Roum1212/todo/pkg/gen/reminder/v1"
)

func TestDeleteReminderRPC_DeleteReminder(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminderID := reminder_id_model.GenerateReminderID()

	request := reminder_v1.DeleteReminderRequest{
		Id: int64(reminderID),
	}

	commandHandlerMock := delete_reminder_command_mock.NewCommandHandlerMock(mc).
		HandleCommandMock.
		Expect(minimock.AnyContext, delete_reminder_command.NewCommand(reminderID)).
		Return(nil)

	deleteReminderPRC := NewDeleteReminderRPC(commandHandlerMock)

	deleteReminderResponse, err := deleteReminderPRC.DeleteReminder(t.Context(), &request)
	require.NoError(t, err)
	require.NotNil(t, deleteReminderResponse)
}

func TestDeleteReminderRPC_DeleteReminder_ErrReminderNotFound(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminderID := reminder_id_model.GenerateReminderID()

	request := reminder_v1.DeleteReminderRequest{
		Id: int64(reminderID),
	}

	commandHandlerMock := delete_reminder_command_mock.NewCommandHandlerMock(mc).
		HandleCommandMock.
		Expect(minimock.AnyContext, delete_reminder_command.NewCommand(reminderID)).
		Return(delete_reminder_command.ErrReminderNotFound)

	deleteReminderPRC := NewDeleteReminderRPC(commandHandlerMock)

	deleteReminderResponse, err := deleteReminderPRC.DeleteReminder(t.Context(), &request)
	require.Error(t, err)
	require.Nil(t, deleteReminderResponse)

	pbStatus, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.NotFound, pbStatus.Code())
}

func TestDeleteReminderRPC_DeleteReminder_Internal(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminderID := reminder_id_model.GenerateReminderID()

	request := reminder_v1.DeleteReminderRequest{
		Id: int64(reminderID),
	}

	commandHandlerMock := delete_reminder_command_mock.NewCommandHandlerMock(mc).
		HandleCommandMock.
		Expect(minimock.AnyContext, delete_reminder_command.NewCommand(reminderID)).
		Return(assert.AnError)

	deleteReminderPRC := NewDeleteReminderRPC(commandHandlerMock)

	deleteReminderResponse, err := deleteReminderPRC.DeleteReminder(t.Context(), &request)
	require.Error(t, err)
	require.Nil(t, deleteReminderResponse)

	pbStatus, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.Internal, pbStatus.Code())
}

func TestDeleteReminderRPC_DeleteReminder_InvalidArgument(t *testing.T) {
	t.Parallel()

	deleteReminderRPC := NewDeleteReminderRPC(nil)

	deleteReminderResponse, err := deleteReminderRPC.DeleteReminder(t.Context(), &reminder_v1.DeleteReminderRequest{})
	require.Error(t, err)
	require.Nil(t, deleteReminderResponse)
}
