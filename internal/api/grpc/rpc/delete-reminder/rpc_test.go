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

	request := reminder_v1.DeleteReminderRequest{
		Id: int64(reminder_id_model.GenerateReminderID()),
	}

	commandHandlerMock := delete_reminder_command_mock.NewCommandHandlerMock(mc).
		HandleCommandMock.
		Expect(t.Context(), delete_reminder_command.NewCommand(reminder_id_model.ReminderID(request.Id))).
		Return(nil)

	deleteReminderPRC := NewDeleteReminderRPC(commandHandlerMock)

	deleteReminderResponse, err := deleteReminderPRC.DeleteReminder(t.Context(), &request)
	require.NoError(t, err)
	require.NotNil(t, deleteReminderResponse)
}

func TestDeleteReminderRPC_DeleteReminder_ErrReminderNotFound(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	request := reminder_v1.DeleteReminderRequest{
		Id: int64(reminder_id_model.GenerateReminderID()),
	}

	commandHandlerMock := delete_reminder_command_mock.NewCommandHandlerMock(mc).
		HandleCommandMock.
		Expect(t.Context(), delete_reminder_command.NewCommand(reminder_id_model.ReminderID(request.Id))).
		Return(delete_reminder_command.ErrReminderNotFound)

	deleteReminderPRC := NewDeleteReminderRPC(commandHandlerMock)

	deleteReminderResponse, err := deleteReminderPRC.DeleteReminder(t.Context(), &request)
	require.Nil(t, deleteReminderResponse)

	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.NotFound, st.Code())
}

func TestDeleteReminderRPC_DeleteReminder_Internal(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	request := reminder_v1.DeleteReminderRequest{
		Id: int64(reminder_id_model.GenerateReminderID()),
	}

	commandHandlerMock := delete_reminder_command_mock.NewCommandHandlerMock(mc).
		HandleCommandMock.
		Expect(t.Context(), delete_reminder_command.NewCommand(reminder_id_model.ReminderID(request.Id))).
		Return(assert.AnError)

	deleteReminderPRC := NewDeleteReminderRPC(commandHandlerMock)

	deleteReminderResponse, err := deleteReminderPRC.DeleteReminder(t.Context(), &request)
	require.Nil(t, deleteReminderResponse)

	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.Internal, st.Code())
}

func TestDeleteReminderRPC_DeleteReminder_InvalidArgument(t *testing.T) {
	t.Parallel()

	request := reminder_v1.DeleteReminderRequest{
		Id: 0,
	}

	deleteReminderRPC := NewDeleteReminderRPC(nil)

	deleteReminderResponse, err := deleteReminderRPC.DeleteReminder(t.Context(), &request)
	require.Nil(t, deleteReminderResponse)

	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.InvalidArgument, st.Code())
}
