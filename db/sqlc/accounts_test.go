package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Silverpoision/simple_bank/util"
	"github.com/stretchr/testify/require"
)

func createAccountRandom(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	account, err := testqueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotEmpty(t, account.CreatedAt)
	require.NotEmpty(t, account.ID)

	return account
}

func TestCreateAcc(t *testing.T) {
	createAccountRandom(t)
}

func TestGetAccount(t *testing.T) {
	account := createAccountRandom(t)
	account2, err := testqueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, account.Currency, account2.Currency)
	require.Equal(t, account.Balance, account2.Balance)
	require.Equal(t, account.Owner, account2.Owner)
	require.WithinDuration(t, account.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account := createAccountRandom(t)
	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomBalance(),
	}

	account2, err := testqueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, account.Currency, account2.Currency)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account.Owner, account2.Owner)
	require.WithinDuration(t, account.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account := createAccountRandom(t)

	err := testqueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	account2, err := testqueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createAccountRandom(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	account, err := testqueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	for _, acc := range account {
		require.NotEmpty(t, acc)
	}
}
