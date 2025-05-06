-- name: CreateFeed :execresult
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id) VALUES (?,?,?,?,?,?);