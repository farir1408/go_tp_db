package helpers

const SelectThreadIdForumSlug = `SELECT id, forum FROM thread
								WHERE slug = $1`

const SelectThreadIdForumSlugByID = `SELECT id, forum FROM thread
									WHERE id = $1`

const SelectThreadID = `SELECT parentId FROM post WHERE id = $1 AND thread = $2`

const CreatePost = `INSERT INTO post (author, created, forum, message, parent, thread)
					VALUES ($1, $2, $3, $4, $5, $6)
					RETURNING id`

const CreatePostParent = `UPDATE post SET parentId = $2, root = $3
							WHERE id = $1`

const SelectPost = `SELECT author, (created AT TIME ZONE 'UTC'), forum, message, parent, thread, isEdited
					FROM post WHERE id = $1`

const UpdatePost = `UPDATE post SET message = coalesce(coalesce(nullif($1, ''), message)),
					isEdited = ('' IS DISTINCT FROM $1)  
					WHERE id = $2
					RETURNING author, (created AT TIME ZONE 'UTC'), forum, id, isEdited, message, parent, thread`

const SelectPostMessage = `SELECT author, (created AT TIME ZONE 'UTC'), forum, id, isEdited, message, parent, thread FROM post WHERE id = $1`

const SelectUserNick = `SELECT about, email, fullname, nickname FROM users WHERE nickname = $1`

const InsertForumUsers = 	`INSERT INTO forum_users (about, email, fullname,
							nickname, forum_slug) VALUES ($1, $2, $3, $4, $5)`