package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type TransferTxParams struct {
	Player   Player `json:"player"`
	FromTeam Team   `json:"from_team"`
	ToTeam   Team   `json:"to_team"`
}

type TransferTxResult struct {
	Transfer Transfer `json:"transfer"`
	Player   Player   `json:"player"`
	FromTeam Team     `json:"from_team"`
	ToTeam   Team     `json:"to_team"`
}

func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		updateArg := UpdatePlayersTeamParams{
			ID:   arg.Player.ID,
			Team: sql.NullInt64{Int64: arg.ToTeam.ID},
		}

		result.FromTeam = arg.FromTeam
		result.ToTeam = arg.ToTeam
		result.Player, err = q.UpdatePlayersTeam(ctx, updateArg)
		if err != nil {
			return err
		}

		transferArg := CreateTransferParams{
			Player:   result.Player.ID,
			FromTeam: result.FromTeam.ID,
			ToTeam:   result.ToTeam.ID,
		}
		result.Transfer, err = q.CreateTransfer(ctx, transferArg)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
