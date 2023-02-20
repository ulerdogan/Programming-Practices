package db

import (
	"basketball_db/utils"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomPlayer(t *testing.T) Player {
	arg := CreatePlayerParams{
		Name: utils.GetRandomString(5),
		Role: Roles(utils.GetRandomRole()),
	}

	player, err := testQueries.CreatePlayer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, player)
	require.Equal(t, player.Name, arg.Name)
	require.Equal(t, player.Role, arg.Role)
	require.Zero(t, player.Team)
	require.NotZero(t, player.CreatedAt)

	return player
}

func TestCreatePlayer(t *testing.T) {
	player := createRandomPlayer(t)
	require.NotEmpty(t, player)
}

func TestGetPlayer(t *testing.T) {
	p := createRandomPlayer(t)

	pt, err := testQueries.GetPlayer(context.Background(), p.ID)
	require.NoError(t, err)
	require.NotEmpty(t, pt)
	require.Equal(t, p.ID, pt.ID)
	require.Equal(t, p.Name, pt.Name)
	require.Equal(t, p.Role, pt.Role)
	require.Equal(t, p.Team, pt.Team)
	require.Equal(t, p.CreatedAt, pt.CreatedAt)
}

func TestUpdatePlayersTeam(t *testing.T) {
	c := createRandomCoach(t)
	team, err := testQueries.CreateTeam(context.Background(), c.ID)
	require.NoError(t, err)

	p := createRandomPlayer(t)
	require.NotEmpty(t, p)

	updateArgs := UpdatePlayersTeamParams{
		ID:   p.ID,
		Team: sql.NullInt64{Int64: team.ID, Valid: true},
	}
	pu, err := testQueries.UpdatePlayersTeam(context.Background(), updateArgs)

	require.NoError(t, err)
	require.NotEmpty(t, pu)
	require.Equal(t, team.ID, pu.Team.Int64)
}

func TestDeletePlayer(t *testing.T) {
	p := createRandomPlayer(t)

	err := testQueries.DeletePlayer(context.Background(), p.ID)
	require.NoError(t, err)

	pt, err := testQueries.GetPlayer(context.Background(), p.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, pt)
}
