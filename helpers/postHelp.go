package helpers

const SelectThreadIdForumSlug = `SELECT id, forum FROM thread
								WHERE slug = $1`

const SelectThreadIdForumSlugByID = `SELECT id, forum FROM thread
									WHERE id = $1`

const SelectThreadID = `SELECT parentId FROM post WHERE id = $1 AND thread = $2`

const CreatePost = `INSERT INTO post (author, created, forum, message, parent, thread)
					VALUES ((SELECT nickname FROM users WHERE nickname = $1),
					$2, $3, $4, $5, $6)
					RETURNING id`

const CreatePostParent = `UPDATE post SET parentId = $2
							WHERE id = $1`

//(coalesce(coalesce(nullif(parent, 0), id)))

const SelectPost = `SELECT author, (created AT TIME ZONE 'UTC'), forum, message::TEXT, parent, thread, isEdited
					FROM post WHERE id = $1`

const UpdatePost = `UPDATE post SET message = coalesce(coalesce(nullif($1, ''), message)),
					isEdited = ('' IS DISTINCT FROM $1)  
					WHERE id = $2
					RETURNING author, (created AT TIME ZONE 'UTC'), forum, id, isEdited, message, parent, thread`

const SelectPostMessage = `SELECT author, (created AT TIME ZONE 'UTC'), forum, id, isEdited, message, parent, thread FROM post WHERE id = $1`