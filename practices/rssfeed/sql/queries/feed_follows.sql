-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN feeds ON inserted_feed_follow.feed_id = feeds.id
INNER JOIN users ON inserted_feed_follow.user_id = users.id;
--

-- name: GetFeedFollowsForUser :many
SELECT feed_follows.*, feeds.name AS feed_name, users.name AS user_name
FROM feed_follows
INNER JOIN feeds ON feed_follows.feed_id = feeds.id
INNER JOIN users ON feeds.user_id = users.id
WHERE feed_follows.user_id = $1;
--


-- name: DeleteFeedFollowByUserAndURL :exec
DELETE FROM feed_follows
USING feeds
WHERE feed_follows.feed_id = feeds.id
AND feed_follows.user_id = $1
AND feeds.url = $2;


-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = now(), updated_at = now()
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT id, name, url, last_fetched_at, user_id, created_at, updated_at
FROM feeds
ORDER BY last_fetched_at NULLS FIRST, last_fetched_at ASC
LIMIT 1;

