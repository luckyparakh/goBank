package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/luckyparakh/goBank/utils"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(utils.GenerateRandomString(32))
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
func TestExpiredJWTToken(t *testing.T) {
	maker, err := NewJWTMaker(utils.GenerateRandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(utils.GenerateRandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.Equal(t, err.Error(), ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestJWTTokenWithAlgoNone(t *testing.T) {
	payload, err := NewPayload(utils.GenerateRandomOwner(), time.Minute)
	require.NoError(t, err)

	jwttoken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwttoken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJWTMaker(utils.GenerateRandomString(32))
	require.NoError(t, err)
	pl, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.Equal(t, err.Error(), ErrInvalidToken.Error())
	require.Nil(t, pl)
}
