package helpers

const SelectThreadIdForumSlug = `SELECT id, forum FROM thread
								WHERE slug = $1`

const SelectThreadIdForumSlugByID = `SELECT id, forum FROM thread
									WHERE id = $1`

const SelectForumID = `SELECT id FROM forum WHERE slug = $1`