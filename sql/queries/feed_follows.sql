-- name: CreateFeedFollow :one 
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT 
    ff.*, 
    f.name AS feed_name,
    u.name AS user_name 
FROM inserted_feed_follow AS ff
INNER JOIN users AS u ON ff.user_id = u.id 
INNER JOIN feeds AS f ON ff.feed_id = f.id;

-- name: GetFeedFollowsForUser :many
SELECT 
    ff.*,
    u.name as user_name,
    f.name as feed_name
from feed_follows AS ff 
INNER JOIN users AS u ON ff.user_id = u.id 
INNER JOIN feeds AS f ON ff.feed_id = f.id 
WHERE u.id = $1;
