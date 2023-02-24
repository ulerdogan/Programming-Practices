package db

import (
	"basketball_db/utils"
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomCoach(t *testing.T) Coach {
	arg := CreateCoachParams{
		Username:       utils.GetRandomString(10),
		HashedPassword: utils.GetRandomString(20),
	}

	coach, err := testQueries.CreateCoach(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, coach)
	require.Equal(t, coach.Username, arg.Username)
	require.NotZero(t, coach.CreatedAt)

	return coach
}

func TestCreateCoach(t *testing.T) {
	coach := createRandomCoach(t)
	require.NotEmpty(t, coach)
}

func TestGetCoachByUsername(t *testing.T) {
	c := createRandomCoach(t)

	ct, err := testQueries.GetCoachByUsername(context.Background(), c.Username)
	require.NoError(t, err)
	require.NotEmpty(t, ct)
	require.Equal(t, c.ID, ct.ID)
	require.Equal(t, c.Username, ct.Username)
	require.Equal(t, c.CreatedAt, ct.CreatedAt)
	require.Equal(t, c.HashedPassword, ct.HashedPassword)
}

func TestUpdateCoachPassword(t *testing.T) {
	c := createRandomCoach(t)
	arg := UpdateCoachPasswordParams{
		Username:       c.Username,
		HashedPassword: utils.GetRandomString(20),
	}

	ct, err := testQueries.UpdateCoachPassword(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, ct)
	require.Equal(t, c.ID, ct.ID)
	require.Equal(t, c.Username, ct.Username)
	require.Equal(t, c.CreatedAt, ct.CreatedAt)
	require.Equal(t, arg.HashedPassword, ct.HashedPassword)

}

func TestDeleteCoach(t *testing.T) {
	c := createRandomCoach(t)
	err := testQueries.DeleteCoach(context.Background(), c.Username)
	require.NoError(t, err)

	ct, err := testQueries.GetCoachByUsername(context.Background(), c.Username)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ct)
}
