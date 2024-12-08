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

func (store *Store) CreateAccountTx(ctx context.Context, arg CreateUserParams) (User, error) {
	var result User
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result, err = q.CreateUser(ctx, arg)
		if err != nil {
			return err
		}
		return nil
	})

	return result, err

}

type LikePlaylistTxParams struct {
	UserID     int32 `json:"User"`
	PlaylistID int32 `json:"PlaylistID"`
}

type LikePlaylistTXResult struct {
	Playlist      Playlist      `json:"Playlist"`
	LikedPlaylist LikedPlaylist `json:"LikedPlaylist"`
}

func (store *Store) LikePlaylistTx(ctx context.Context, arg LikePlaylistTxParams) (LikePlaylistTXResult, error) {
	var result LikePlaylistTXResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.LikedPlaylist, err = q.AddUserLike(ctx, AddUserLikeParams{
			PlaylistID: arg.PlaylistID,
			UserID:     arg.UserID,
		})
		if err != nil {
			return err
		}

		result.Playlist, err = q.LikePlaylist(ctx, arg.PlaylistID)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err

}
