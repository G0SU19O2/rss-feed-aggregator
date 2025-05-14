-- name: CreatePost :exec
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id) VALUES (?,?,?,?,?,?,?, ?);

-- name: GetPosts :many
SELECT * FROM posts ORDER BY created_at DESC LIMIT ?;

-- name: GetPostsForUser :many
SELECT p.* FROM posts p
JOIN feeds f ON p.feed_id = f.id
JOIN feed_follows fw ON f.id = fw.feed_id
WHERE fw.user_id = ?
ORDER BY p.published_at DESC
LIMIT ?;
