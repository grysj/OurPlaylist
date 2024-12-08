// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: liked_playlist.sql

package db

import (
	"context"
)

const addUserLike = `-- name: AddUserLike :one
INSERT INTO liked_playlist (playlist_id, user_id)
VALUES ($1, $2)
RETURNING id, playlist_id, user_id
`

type AddUserLikeParams struct {
	PlaylistID int32 `json:"playlist_id"`
	UserID     int32 `json:"user_id"`
}

func (q *Queries) AddUserLike(ctx context.Context, arg AddUserLikeParams) (LikedPlaylist, error) {
	row := q.db.QueryRowContext(ctx, addUserLike, arg.PlaylistID, arg.UserID)
	var i LikedPlaylist
	err := row.Scan(&i.ID, &i.PlaylistID, &i.UserID)
	return i, err
}

const countLikesForPlaylist = `-- name: CountLikesForPlaylist :one
SELECT COUNT(*) AS like_count 
FROM liked_playlist 
WHERE playlist_id = $1
`

func (q *Queries) CountLikesForPlaylist(ctx context.Context, playlistID int32) (int64, error) {
	row := q.db.QueryRowContext(ctx, countLikesForPlaylist, playlistID)
	var like_count int64
	err := row.Scan(&like_count)
	return like_count, err
}

const deleteUserLike = `-- name: DeleteUserLike :exec
DELETE FROM liked_playlist WHERE id = $1
`

func (q *Queries) DeleteUserLike(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUserLike, id)
	return err
}

const getLikedPlaylistsByUser = `-- name: GetLikedPlaylistsByUser :many
SELECT playlists.id, playlists.user_id, playlists.created_at, playlists.link, playlists.like_count, playlists.description 
FROM liked_playlist 
JOIN playlists ON liked_playlist.playlist_id = playlists.id 
WHERE liked_playlist.user_id = $1
`

func (q *Queries) GetLikedPlaylistsByUser(ctx context.Context, userID int32) ([]Playlist, error) {
	rows, err := q.db.QueryContext(ctx, getLikedPlaylistsByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Playlist{}
	for rows.Next() {
		var i Playlist
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CreatedAt,
			&i.Link,
			&i.LikeCount,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}