package helpers

const SelectForumCreate = `SELECT posts, slug::TEXT, threads, title, author FROM forum
						   WHERE slug = $1 AND author = $2`

const SelectForumByUser = `INSERT INTO forum (slug, title, author) VALUES($1, $2, (SELECT nickname FROM users
						   WHERE lower(nickname) = lower($3)))
						   RETURNING author::TEXT`

const SelectForumDetail = `SELECT posts, slug, threads, title, author FROM forum
							WHERE slug = $1`

const SelectThreadCreate = `SELECT author, created, forum, id, message, slug, title, votes
							FROM thread WHERE message = $1 
							AND title = $2 AND lower(slug) = lower($3)`

const SelectThreadByUser = `INSERT INTO thread (author, message, title, forum, slug, created)
							VALUES((SELECT nickname FROM users WHERE nickname = $1),
							$2, $3, (SELECT slug FROM forum WHERE slug = $4), $5, $6)
							RETURNING id`

const SelectThreadsDesc = `SELECT author, created, forum, id, message, slug, title FROM thread
					   WHERE forum = (SELECT slug FROM forum WHERE slug = $1)
					   ORDER By created DESC
					   LIMIT $2::TEXT::INTEGER`

const SelectThreads = `SELECT author, created, forum, id, message, slug, title FROM thread
					   WHERE forum = (SELECT slug FROM forum WHERE slug = $1)
					   ORDER By created
					   LIMIT $2::TEXT::INTEGER`

const SelectThreadsSince = `SELECT author, created, forum, id, message, slug, title FROM thread
					   WHERE forum = (SELECT slug FROM forum WHERE slug = $1)
					   AND created >= $2::TEXT::TIMESTAMPTZ
					   ORDER By created
					   LIMIT $3::TEXT::INTEGER`

const SelectThreadsSinceDesc = `SELECT author, created, forum, id, message, slug, title FROM thread
					   WHERE forum = (SELECT slug FROM forum WHERE slug = $1)
					   AND created <= $2::TEXT::TIMESTAMPTZ
					   ORDER By created DESC
					   LIMIT $3::TEXT::INTEGER`
