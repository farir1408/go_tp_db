package helpers

const SelectForumCreate = `SELECT posts, slug, threads, title, author FROM forum
					WHERE slug = $1 AND title = $2 AND author = $3`

const SelectForumByUser = `INSERT INTO forum (slug, title, author) VALUES($1, $2, $3)`


