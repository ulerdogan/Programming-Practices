package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	c := createRandomCoach(t)
	team, err := testQueries.CreateTeam(context.Background(), c.ID)
	require.NoError(t, err)

	p1 := createRandomPlayer(t)

	updateArgs := UpdatePlayersTeamParams{
		ID:   p1.ID,
		Team: sql.NullInt64{Int64: team.ID, Valid: true},
	}
	p1u, err := store.UpdatePlayersTeam(context.Background(), updateArgs)
	require.NoError(t, err)
	require.NotEmpty(t, p1u)
	require.Empty(t, p1.Team)
	require.Equal(t, team.ID, p1u.Team.Int64)
}
