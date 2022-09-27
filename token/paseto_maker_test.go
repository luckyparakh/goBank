package token

import (
	"testing"
	"time"

	"github.com/luckyparakh/goBank/utils"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(utils.GenerateRandomString(32))
	require.NoError(t, err)
	username := utils.GenerateRandomOwner()
	duration := time.Minute
	issuedAt := time.Now()
	exipredAt := time.Now().Add(duration)
	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.NotEmpty(t, payload.Id)
	require.Equal(t, payload.Username, username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, exipredAt, payload.ExpiredAt, time.Second)

}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(utils.GenerateRandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(utils.GenerateRandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.Equal(t, err.Error(), ErrExpiredToken.Error())
	require.Nil(t, payload)
}