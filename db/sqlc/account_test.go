package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/luckyparakh/goBank/utils"
	"github.com/stretchr/testify/require"
)

func randomAccCreation(t *testing.T) Account {
	user := randomUserCreation(t)
	args := CreateAccountParams{
		Owner:    user.Username,
		Balance:  utils.GenerateRandomMoney(),
		Currency: utils.GenerateRandomCurrency(),
	}
	acc, err := testQueries.CreateAccount(context.Background(), args)
	require.NotEmpty(t, acc)
	require.NoError(t, err)
	require.Equal(t, acc.Balance, args.Balance)
	require.Equal(t, acc.Owner, args.Owner)
	require.Equal(t, acc.Currency, args.Currency)
	require.NotZero(t, acc.ID)
	require.NotZero(t, acc.CreatedAt)
	return acc
}
func TestCreateAccount(t *testing.T) {
	randomAccCreation(t)
}

func TestGetAccount(t *testing.T) {
	acc1 := randomAccCreation(t)
	acc2, err := testQueries.GetAccount(context.Background(), acc1.ID)
	require.NotEmpty(t, acc2)
	require.Empty(t, err)
	require.Equal(t, acc1.ID, acc2.ID)
	require.WithinDuration(t, acc1.CreatedAt, acc2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	acc1 := randomAccCreation(t)
	args := UpdateAccountParams{
		ID:      acc1.ID,
		Balance: utils.GenerateRandomMoney(),
	}
	acc2, err := testQueries.UpdateAccount(context.Background(), args)
	require.NotEmpty(t, acc2)
	require.Empty(t, err)
	require.Empty(t, err)
	require.Equal(t, acc1.ID, acc2.ID)
	require.Equal(t, args.Balance, acc2.Balance)
	require.WithinDuration(t, acc1.CreatedAt, acc2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	acc1 := randomAccCreation(t)
	err := testQueries.DeleteAccount(context.Background(), acc1.ID)
	require.Empty(t, err)

	acc, err := testQueries.GetAccount(context.Background(), acc1.ID)
	require.Empty(t, acc)
	require.NotEmpty(t, err)
	require.Equal(t, err.Error(), sql.ErrNoRows.Error())
}

func TestListAccounts(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = randomAccCreation(t)
	}
	args := ListAccountsParams{
		Limit:  5,
		Offset: 0,
		Owner:  lastAccount.Owner,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), args)
	require.NotEmpty(t, accounts)
	require.Empty(t, err)
}
