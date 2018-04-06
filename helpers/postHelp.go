package helpers

const SelectThreadIdForumSlug = `SELECT id, forum FROM thread
								WHERE slug = $1`

const SelectThreadIdForumSlugByID = `SELECT id, forum FROM thread
									WHERE id = $1`

const SelectThreadID = `SELECT id FROM thread WHERE id = $1`

const CreatePost = `INSERT INTO post (author, created, forum, message, parent, thread)
					VALUES ((SELECT nickname FROM users WHERE nickname = $1),
					$2, $3, $4, $5, $6)
					RETURNING id`

//const SelectPost = `SELECT author, created, forum, id, message, slug, title FROM post
//					WHERE slug = $1`