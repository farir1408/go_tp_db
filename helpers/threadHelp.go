package helpers

const SelectThreadBySlug = `SELECT id, author, created, forum, message, slug,
							title, votes FROM thread
							WHERE slug = $1`

const SelectThreadById = `SELECT id, author, created, forum, message, slug,
							title, votes FROM thread
							WHERE id = $1`