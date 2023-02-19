package token

import (
	"testing"
	"time"

	"beebeewijaya.com/util"
	"github.com/stretchr/testify/require"
)

func createToken(t *testing.T) (*JWTMaker, string) {
	secretKey := util.RandomString(16)
	username := util.RandomInt(0, 1000)
	email := util.RandomEmail()
	duration := time.Minute * 5

	maker := &JWTMaker{
		secretKey: secretKey,
	}

	token, err := maker.GenerateToken(username, email, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	return maker, token
}

func TestCreateToken(t *testing.T) {
	createToken(t)
}

func TestVerifyToken(t *testing.T) {
	maker, token := createToken(t)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.NotZero(t, payload.IssuedAt)
	require.NotZero(t, payload.ExpiredAt)
	require.NotEmpty(t, payload.Username)
}
