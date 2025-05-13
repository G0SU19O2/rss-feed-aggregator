-- name: CreateFeed :execresult
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id) VALUES (?,?,?,?,?,?);

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedByURL :one
SELECT * FROM feeds WHERE url = ? LIMIT 1;

-- name: MarkFeedFetched :exec
UPDATE feeds SET last_fetched_at = ?, updated_at = ? WHERE id = ?;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY (last_fetched_at IS NOT NULL), last_fetched_at ASC, updated_at ASC
LIMIT 1;

-- name: DeleteFeed :exec

DELETE FROM feeds WHERE id = ?;