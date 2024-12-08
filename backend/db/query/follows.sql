-- name: AddFollow :exec
INSERT INTO follows (follower_id, follows_id) VALUES ($1, $2);


-- name: RemoveFollow :exec
DELETE FROM follows WHERE follower_id = $1 AND follows_id = $2;