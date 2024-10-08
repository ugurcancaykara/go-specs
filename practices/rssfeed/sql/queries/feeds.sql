-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetAuthor :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetFeedByURL :one
SELECT id, created_at, updated_at, name, url, user_id
FROM feeds
WHERE url = $1
LIMIT 1;
