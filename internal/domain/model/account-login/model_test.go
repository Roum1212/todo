package account_login_model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewAccountLogin(t *testing.T) {
	t.Parallel()

	login := NewAccountLogin("abc123")
	require.Equal(t, "abc123", string(login))
}
