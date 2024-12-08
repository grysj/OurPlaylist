-- name: CountLikesForPlaylist :one
SELECT COUNT(*) AS like_count 
FROM liked_playlist 
WHERE playlist_id = $1;


-- name: GetLikedPlaylistsByUser :many
SELECT playlists.* 
FROM liked_playlist 
JOIN playlists ON liked_playlist.playlist_id = playlists.id 
WHERE liked_playlist.user_id = $1;


-- name: AddUserLike :one
INSERT INTO liked_playlist (playlist_id, user_id)
VALUES ($1, $2)
RETURNING *;


-- name: DeleteUserLike :exec
DELETE FROM liked_playlist WHERE id = $1;