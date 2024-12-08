-- name: GetPlaylistsByUserID :many
SELECT * FROM playlists WHERE user_id = $1 ORDER BY created_at DESC;

-- name: CreatePlaylist :one
INSERT INTO playlists (user_id, link, description) VALUES ($1, $2, $3)
RETURNING *;

-- name: DeletePlaylist :exec
DELETE FROM playlists WHERE id = $1;


-- name: LikePlaylist :one
UPDATE playlists
SET like_count = like_count + 1
WHERE id = $1
RETURNING *;

-- name: UnlikePlaylist :one
UPDATE playlists
SET like_count = like_count - 1
WHERE id = $1
RETURNING *;