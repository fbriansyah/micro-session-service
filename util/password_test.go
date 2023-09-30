package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGeneratePassword(t *testing.T) {
	pass := "S3cr3t"
	hasedPassword, err := HashPassword(pass)
	require.NoError(t, err)

	err = CheckPassword(pass, hasedPassword)
	require.NoError(t, err)
}
