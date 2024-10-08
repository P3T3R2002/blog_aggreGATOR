-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT 
    iff.* , 
    f.name AS feed_name,
    u.name aS user_name 
FROM inserted_feed_follow iff
INNER JOIN users u ON u.id = iff.user_id
INNER JOIN feeds f ON f.id = iff.feed_id;

-- name: GetFeedFollowsForUser :many
SELECT 
    ff.* , 
    f.name AS feed_name,
    u.name aS user_name
FROM feed_follows ff
INNER JOIN users u ON u.id = ff.user_id
INNER JOIN feeds f ON f.id = ff.feed_id
WHERE u.id = $1;

-- name: FeedUnfollow :exec
DELETE FROM feed_follows
WHERE $1 = user_id AND $2 = feed_id;


