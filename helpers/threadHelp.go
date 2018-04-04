package helpers

const SelectThreadBySlug = `SELECT id, author, created, forum, message, slug,
							title, votes FROM thread
							WHERE slug = $1`

const SelectThreadById = `SELECT id, author, created, forum, message, slug,
							title, votes FROM thread
							WHERE id = $1`

const UpdateThreadBySlug = `UPDATE thread SET 
						message = coalesce(coalesce(nullif($1, ''), message)),
						title = coalesce(coalesce(nullif($2, ''), title))
					  WHERE slug = $3`

const UpdateThreadById = `UPDATE thread SET 
						message = coalesce(coalesce(nullif($1, ''), message)),
						title = coalesce(coalesce(nullif($2, ''), title))
					  WHERE id = $3`
