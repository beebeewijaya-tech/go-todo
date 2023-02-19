package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	randomPassword := RandomString(8)

	hashed, err := HashPassword(randomPassword)
	require.NoError(t, err)
	require.NotEmpty(t, hashed)

	// Correct Password
	errCorrectPassword := CheckPassword(randomPassword, hashed)
	require.NoError(t, errCorrectPassword)

	// False Password
	errFalsePassword := CheckPassword(RandomString(8), hashed)
	require.Error(t, errFalsePassword)
}
