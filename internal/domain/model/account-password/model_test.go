package account_password_model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewAccountPassword(t *testing.T) {
	t.Parallel()

	password := NewAccountPassword("abc123")
	require.Equal(t, "abc123", string(password))
}
