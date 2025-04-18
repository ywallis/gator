-- name: CreatePost :one

INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
	$1,
	$2,
	$3,
	$4,
	$5,
	$6,
	$7,
	$8
	)
RETURNING *;


-- name: GetPostsForUser :many
SELECT posts.* FROM posts 
WHERE feed_id IN (
	SELECT feed_follows.feed_id FROM feed_follows
	INNER JOIN users ON feed_follows.user_id = users.id 
	INNER JOIN feeds ON feed_follows.feed_id = feeds.id 
	WHERE users.id = $1
)
ORDER BY published_at DESC NULLS LAST
LIMIT $2
;
