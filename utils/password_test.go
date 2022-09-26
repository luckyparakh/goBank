package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := GenerateRandomString(10)
	hp, err := HashedPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hp)
	err = CheckPassword(password, hp)
	require.NoError(t, err)

	wrong_password := GenerateRandomString(10)
	err = CheckPassword(wrong_password, hp)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hp1, err := HashedPassword(password)
	require.NoError(t, err)
	require.NotEqual(t, hp, hp1)
}
