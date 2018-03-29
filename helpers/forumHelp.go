package helpers

const SelectForumCreate = `SELECT posts, slug, threads, title, author FROM forum
					WHERE slug = $1 AND title = $2 AND author = $3`

const SelectForumByUser = `INSERT INTO forum (slug, title, author) VALUES($1, $2, $3)`

const SelectForumDetail = `SELECT posts, slug, threads, title, author FROM forum
					WHERE slug = $1`

const SelectThreadCreate = `SELECT author, created, forum, id, message, slug, title, votes
							FROM thread WHERE author = $1 AND message = $2 AND title = $3 AND created = $4`

const SelectThreadByUser = `INSERT INTO thread (author, message, title, forum, slug, created)
							VALUES($1, $2, $3, $4, $5, $6)`
